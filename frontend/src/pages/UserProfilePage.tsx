// frontend/src/pages/UserProfilePage.tsx
import { useCallback, useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import StatsBar from "../components/StatsBar";
import Modal from "../components/Modal";
import AchievementBadge from "../components/AchievementBadge";
import { languages } from "../components/LanguageIcons";
import { BoltIcon, ChevronIcon, FlameIcon, LeagueShield, MedalIcon } from "../components/ProfileIcons";
import { WEEK_DAYS, getProfile, type Achievement } from "../data/mockData";
import { apiErrorMessage } from "../services/errors";
import { mediaSrc } from "../services/mediaUrl";
import { followUser, getFollowList, getPublicProfile, unfollowUser } from "../services/socialService";
import type { PublicProfile, UserCard } from "../types/social";
import { useMyStats } from "../hooks/useMyStats";
import profileImg from "../assets/profile.png";
import followersImg from "../assets/pengikut.png";
import "./pageShell.css";
import "./UserProfilePage.css";

type Series = { label: string; values: number[]; color: string };

// Pilihan alasan laporan. Daftar resmi nanti ditentukan backend.
const REPORT_REASONS = ["Ketelanjangan", "Spam", "Lainnya"];

function joinedText(iso: string): string {
  const d = new Date(iso);
  if (Number.isNaN(d.getTime())) return "";
  return `Bergabung ${d.toLocaleDateString("id-ID", { month: "long", year: "numeric" })}`;
}

function Avatar({ name, src, className }: { name: string; src: string; className: string }) {
  return src ? (
    <img src={mediaSrc(src)} alt="" className={className} style={{ objectFit: "cover" }} />
  ) : (
    <span className={className}>{name.charAt(0).toUpperCase()}</span>
  );
}

/* ---------- Grafik XP minggu ini (SVG) ---------- */
function XpChart({ series }: { series: Series[] }) {
  const W = 640, H = 280, L = 44, R = 16, T = 12, B = 32;
  const top = Math.max(...series.flatMap((s) => s.values));
  const max = Math.max(80, Math.ceil(top / 20) * 20);
  const x = (i: number) => L + (i * (W - L - R)) / (WEEK_DAYS.length - 1);
  const y = (v: number) => T + (H - T - B) * (1 - v / max);
  const ticks = Array.from({ length: max / 20 + 1 }, (_, i) => i * 20);

  return (
    <svg viewBox={`0 0 ${W} ${H}`} className="xp__svg" role="img" aria-label="Grafik XP minggu ini">
      {ticks.map((t) => (
        <g key={t}>
          <line x1={L} x2={W - R} y1={y(t)} y2={y(t)} stroke="#e5e5e5" strokeWidth="2" />
          <text x={L - 10} y={y(t) + 5} textAnchor="end" fontSize="15" fontWeight="700" fill="#afafaf">
            {t}
          </text>
        </g>
      ))}
      {WEEK_DAYS.map((d, i) => (
        <text key={d + i} x={x(i)} y={H - 8} textAnchor="middle" fontSize="15" fontWeight="700" fill="#afafaf">
          {d}
        </text>
      ))}
      {series.map((s) => (
        <g key={s.label}>
          <polyline
            points={s.values.map((v, i) => `${x(i)},${y(v)}`).join(" ")}
            fill="none"
            stroke={s.color}
            strokeWidth="3"
            strokeLinejoin="round"
          />
          {s.values.map((v, i) => (
            <circle key={i} cx={x(i)} cy={y(v)} r="5" fill={s.color} />
          ))}
        </g>
      ))}
    </svg>
  );
}

/* ---------- Ikon kecil ---------- */
function FlagIcon() {
  return (
    <svg viewBox="0 0 24 24" aria-hidden="true">
      <path d="M5 21V4M5 5h13l-3 4 3 4H5" fill="none" stroke="currentColor" strokeWidth="2.5" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  );
}
function PersonPlusIcon() {
  return (
    <svg viewBox="0 0 24 24" aria-hidden="true">
      <circle cx="9" cy="8" r="4" fill="currentColor" />
      <path d="M2 21c0-4 3-7 7-7s7 3 7 7z" fill="currentColor" />
      <path d="M19 7v6M16 10h6" stroke="currentColor" strokeWidth="2.5" strokeLinecap="round" />
    </svg>
  );
}
function ChevronDownIcon() {
  return (
    <svg viewBox="0 0 24 24" aria-hidden="true">
      <path d="M5 9l7 7 7-7" fill="none" stroke="currentColor" strokeWidth="3" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  );
}

export default function UserProfilePage() {
  const { username = "" } = useParams();
  // Data contoh (statistik, grafik, pencapaian) hanya ada untuk akun demo;
  // identitas, foto, dan follow selalu dari backend.
  const mock = getProfile(username);

  const [real, setReal] = useState<PublicProfile | null>(null);
  const [loading, setLoading] = useState(true);
  const [followBusy, setFollowBusy] = useState(false);
  const [followError, setFollowError] = useState("");
  const [list, setList] = useState<UserCard[]>([]);

  const [tab, setTab] = useState<"following" | "followers">("following");
  const [showLangs, setShowLangs] = useState(false);
  const myStats = useMyStats();
  const [badge, setBadge] = useState<Achievement | null>(null);

  // Laporan: pop-up, alasan terpilih, dan akun yang baru dilaporkan
  const [showReport, setShowReport] = useState(false);
  const [reason, setReason] = useState("");
  const [reportedFor, setReportedFor] = useState("");

  useEffect(() => {
    let cancelled = false;
    setLoading(true);
    setFollowError("");
    getPublicProfile(username)
      .then((p) => {
        if (!cancelled) setReal(p);
      })
      .catch(() => {
        if (!cancelled) setReal(null);
      })
      .finally(() => {
        if (!cancelled) setLoading(false);
      });
    return () => {
      cancelled = true;
    };
  }, [username]);

  // Daftar mengikuti/pengikut dimuat ulang saat tab atau jumlahnya berubah.
  const realName = real?.username;
  const followers = real?.followers_count;
  const following = real?.following_count;
  useEffect(() => {
    if (!realName) return;
    let cancelled = false;
    getFollowList(realName, tab)
      .then((d) => {
        if (!cancelled) setList(d.users);
      })
      .catch(() => {
        if (!cancelled) setList([]);
      });
    return () => {
      cancelled = true;
    };
  }, [realName, tab, followers, following]);

  const toggleFollow = useCallback(
    async (target: string, currentlyFollowing: boolean) => {
      if (followBusy) return;
      setFollowBusy(true);
      setFollowError("");
      try {
        const updated = currentlyFollowing ? await unfollowUser(target) : await followUser(target);
        if (target === username) {
          setReal(updated);
        } else {
          // Baris pada daftar: muat ulang profil halaman ini agar jumlahnya akurat.
          setReal(await getPublicProfile(username));
        }
      } catch (e) {
        setFollowError(apiErrorMessage(e, "Gagal memperbarui status ikuti. Coba lagi."));
      } finally {
        setFollowBusy(false);
      }
    },
    [followBusy, username],
  );

  if (loading) {
    return (
      <div className="shell">
        <div className="shell__main">
          <p role="status">Memuat profil...</p>
        </div>
      </div>
    );
  }

  if (!real) {
    return (
      <div className="shell">
        <div className="shell__main">
          <h1>Pengguna tidak ditemukan</h1>
          <p>
            Akun <b>{username}</b> tidak ada. <Link to="/papan-skor">Kembali ke papan skor</Link>
          </p>
        </div>
      </div>
    );
  }

  const profile = {
    username: real.username,
    displayName: real.name,
    joined: joinedText(real.joined_at),
  };
  const isMe = real.is_me;
  const series: Series[] = mock ? [{ label: profile.username, values: mock.weeklyXp, color: "#1cb0f6" }] : [];
  const shownLangs = mock ? mock.languages.slice(0, 3) : [];
  const extraLangs = mock ? mock.languages.length - shownLangs.length : 0;

  const closeReport = () => {
    setShowReport(false);
    setReason("");
  };

  const submitReport = () => {
    if (!reason) return;
    // TODO: kirim ke backend setelah endpoint laporan tersedia (alasan divalidasi di backend).
    setReportedFor(profile.username);
    closeReport();
  };

  return (
    <div className="shell">
      <div className="shell__main uprofile">
        <div
          className="uprofile__banner"
          style={
            real.banner_url
              ? { backgroundImage: `url("${mediaSrc(real.banner_url)}")`, backgroundSize: "cover", backgroundPosition: "center" }
              : undefined
          }
        >
          {real.avatar_url ? (
            <img src={mediaSrc(real.avatar_url)} alt="" className="uprofile__photo" style={{ borderRadius: "50%", objectFit: "cover" }} />
          ) : (
            <img src={profileImg} alt="" className="uprofile__photo" />
          )}
        </div>

        <h1 className="uprofile__name">{profile.displayName}</h1>
        <p className="uprofile__user">{profile.username}</p>
        <p className="uprofile__joined">{profile.joined}</p>

        <div className="uprofile__meta">
          <div className="uprofile__follow">
            <span>mengikuti {real.following_count}</span>
            <span>diikuti {real.followers_count}</span>
          </div>

          {/* Ikon bahasa pemrograman yang dipelajari (pengganti bendera). Klik untuk melihat daftarnya. */}
          {mock && (
          <button
            type="button"
            className="uprofile__langs"
            aria-label="Lihat bahasa yang dipelajari"
            onClick={() => setShowLangs(true)}
          >
            {shownLangs.map((l) => {
              const lang = languages.find((x) => x.id === l.id);
              return lang ? (
                <span key={l.id} className="uprofile__chip">
                  <lang.Icon />
                </span>
              ) : null;
            })}
            {extraLangs > 0 && <span className="uprofile__more">+{extraLangs}</span>}
          </button>
          )}
        </div>

        {/* Tombol Laporkan: membuka pop-up pilihan alasan. Tidak tampil di profil sendiri. */}
        {!isMe && (
          <div className="uprofile__actions">
            <button
              type="button"
              className="uprofile__btn"
              disabled={followBusy}
              onClick={() => toggleFollow(real.username, real.is_following)}
            >
              <PersonPlusIcon /> {real.is_following ? "Berhenti mengikuti" : real.follows_you ? "Ikuti balik" : "Ikuti"}
            </button>
            <button type="button" className="uprofile__btn" onClick={() => setShowReport(true)}>
              <FlagIcon /> Laporkan
            </button>
            <button type="button" className="uprofile__btn uprofile__btn--icon" aria-label="Opsi lainnya">
              •••
            </button>
          </div>
        )}
        {followError && (
          <p className="uprofile__note" role="alert">
            {followError}
          </p>
        )}
        {reportedFor === profile.username && (
          <p className="uprofile__note" role="status">
            Laporan dicatat. Catatan: belum dikirim ke server karena endpoint-nya belum tersedia.
          </p>
        )}

        <div className="uprofile__divider" />

        <div className="section-head">
          <h2>Statistik</h2>
        </div>
        <div className="uprofile__stats">
          {mock && (
            <div className="stat-card">
              <FlameIcon on={mock.streak > 0} />
              <div>
                <b>{mock.streak}</b>
                <span>Runtunan hari</span>
              </div>
            </div>
          )}
          <div className="stat-card">
            <BoltIcon />
            <div>
              <b>{real.xp}</b>
              <span>Total XP</span>
            </div>
          </div>
          <div className="stat-card">
            <MedalIcon />
            <div>
              <b>{real.level}</b>
              <span>Level</span>
            </div>
          </div>
          {mock && (
            <div className="stat-card">
              <LeagueShield tone={mock.leagueTone} />
              <div>
                <b>{mock.league}</b>
                <span>Liga saat ini</span>
              </div>
              <em className="stat-card__tag">Minggu ke-{mock.leagueWeek}</em>
            </div>
          )}
        </div>

        {mock && (
          <>

        <div className="section-head">
          <h2>XP minggu ini</h2>
        </div>
        <section className="panel xp">
          <ul className="xp__legend">
            {series.map((s) => (
              <li key={s.label}>
                <span className="xp__dot" style={{ background: s.color }} />
                <span className="xp__name" style={{ color: s.color }}>
                  {s.label}
                </span>
                <b style={{ color: s.color }}>{s.values[s.values.length - 1]} XP</b>
              </li>
            ))}
          </ul>
          <XpChart series={series} />
        </section>

        <div className="section-head">
          <h2>Pencapaian</h2>
          <Link to={`/pengguna/${profile.username}/pencapaian`}>Lihat semua</Link>
        </div>
        {/* Lima pencapaian dalam satu baris */}
        <section className="panel uprofile__badges">
          {mock.achievements.slice(0, 5).map((a) => (
            <button
              key={a.id}
              type="button"
              className="uprofile__badge"
              aria-label={`Lihat pencapaian ${a.name}`}
              onClick={() => setBadge(a)}
            >
              <AchievementBadge icon={a.icon} color={a.color} level={a.level} />
            </button>
          ))}
        </section>
          </>
        )}
      </div>

      <aside className="shell__side">
        {myStats && <StatsBar stats={myStats} />}

        <section className="panel follow">
          <div className="follow__tabs" role="tablist">
            <button
              type="button"
              role="tab"
              aria-selected={tab === "following"}
              className="follow__tab"
              onClick={() => setTab("following")}
            >
              Mengikuti
            </button>
            <button
              type="button"
              role="tab"
              aria-selected={tab === "followers"}
              className="follow__tab"
              onClick={() => setTab("followers")}
            >
              Pengikut
            </button>
          </div>

          {list.length > 0 ? (
            <ul className="follow__list">
              {list.map((f) => (
                <li key={f.username} className="follow-row">
                  <Avatar name={f.name} src={f.avatar_url} className="follow-row__avatar" />
                  <div>
                    <Link to={`/pengguna/${f.username}`}>
                      <b>{f.name}</b>
                    </Link>
                    <span>{f.xp} XP</span>
                  </div>
                  {!f.is_me && (
                    <button
                      type="button"
                      className="uprofile__btn"
                      disabled={followBusy}
                      onClick={() => toggleFollow(f.username, f.is_following)}
                    >
                      {f.is_following ? "Diikuti" : "Ikuti"}
                    </button>
                  )}
                </li>
              ))}
            </ul>
          ) : tab === "following" ? (
            <p className="follow__empty">Belum mengikuti siapa pun</p>
          ) : (
            <div className="follow__box">
              <h3>Jadilah pengikut pertamanya!</h3>
              <img src={followersImg} alt="" className="follow__photo" />
              <p>Belajar lebih seru dan efektif saat kamu terhubung dengan yang lain.</p>
              {!isMe && (
                <button
                  type="button"
                  className="follow__btn"
                  disabled={followBusy}
                  onClick={() => toggleFollow(real.username, real.is_following)}
                >
                  <PersonPlusIcon /> {real.is_following ? "Diikuti" : "Ikuti"}
                </button>
              )}
            </div>
          )}
        </section>
      </aside>

      {/* Daftar bahasa pemrograman yang dipelajari */}
      {showLangs && mock && (
        <Modal title={`Bahasa ${profile.displayName}`} onClose={() => setShowLangs(false)}>
          <ul className="lang-list">
            {mock.languages.map((l) => {
              const lang = languages.find((x) => x.id === l.id);
              if (!lang) return null;
              return (
                <li key={l.id} className="lang-row">
                  <span className="lang-row__icon">
                    <lang.Icon />
                  </span>
                  <b className="lang-row__name">{lang.name}</b>
                  <span className="lang-row__xp">{l.xp} XP</span>
                  <span className="lang-row__go">
                    <ChevronIcon />
                  </span>
                </li>
              );
            })}
          </ul>
        </Modal>
      )}

      {/* Detail pencapaian */}
      {badge && (
        <Modal onClose={() => setBadge(null)}>
          <div className="ach-modal">
            <AchievementBadge icon={badge.icon} color={badge.color} level={badge.level} size="lg" />
            <h2>{badge.name}</h2>
            <p>
              {badge.goal ? (
                <>
                  <b>Level Selanjutnya:</b> {badge.goal}
                </>
              ) : (
                badge.doneText
              )}
            </p>
          </div>
        </Modal>
      )}

      {/* Pop-up laporkan pengguna */}
      {showReport && (
        <Modal onClose={closeReport}>
          <div className="report">
            <h2>Laporkan {profile.username}</h2>

            <div className="report__select">
              <select
                aria-label="Pilih alasan laporan"
                className={reason ? "" : "report__select--empty"}
                value={reason}
                onChange={(e) => setReason(e.target.value)}
              >
                <option value="">Pilih alasan laporan ini...</option>
                {REPORT_REASONS.map((r) => (
                  <option key={r} value={r}>
                    {r}
                  </option>
                ))}
              </select>
              <ChevronDownIcon />
            </div>

            <div className="report__actions">
              <button type="button" className="report__cancel" onClick={closeReport}>
                Tolak
              </button>
              <button type="button" className="report__submit" disabled={!reason} onClick={submitReport}>
                Laporkan
              </button>
            </div>
          </div>
        </Modal>
      )}
    </div>
  );
}