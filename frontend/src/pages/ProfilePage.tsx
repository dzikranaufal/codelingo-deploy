
import { useEffect, useRef, useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import { deleteProfileImage, getMyProfile, uploadProfileImage } from "../services/profileService";
import { mediaSrc } from "../services/mediaUrl";
import { apiErrorMessage } from "../services/errors";
import { getMyStatsDTO, type MyStats } from "../services/statsService";
import type { Profile } from "../types/profile";
import mascot from "../assets/logo.png";
import "./ProfilePage.css";

const show = (value?: number) => (value === undefined ? "—" : value);

// Batas sama dengan backend (internal/media). Backend tetap yang menentukan.
const LIMIT_MB = { avatar: 2, banner: 4 } as const;
const ACCEPT = "image/jpeg,image/png,image/gif,image/webp";

export default function ProfilePage() {
  const navigate = useNavigate();
  const [profile, setProfile] = useState<Profile | null>(null);
  const [stats, setStats] = useState<MyStats | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [mediaBusy, setMediaBusy] = useState(false);
  const [mediaError, setMediaError] = useState("");
  const avatarInput = useRef<HTMLInputElement>(null);
  const bannerInput = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (!localStorage.getItem("token")) {
      navigate("/login", { replace: true });
      return;
    }

    let cancelled = false;

    (async () => {
      try {
        // Streak dan nyawa terkini ada di /me/stats; kalau gagal, profil tetap tampil.
        const [data, statsData] = await Promise.all([
          getMyProfile(),
          getMyStatsDTO().catch(() => null),
        ]);
        if (!cancelled) {
          setProfile(data);
          setStats(statsData);
        }
      } catch (err) {
        if (cancelled) return;

        if (axios.isAxiosError(err)) {
          if (err.response?.status === 401) {
            localStorage.removeItem("token");
            navigate("/login", { replace: true });
            return;
          }

          setError(
            err.response
              ? err.response.data?.error ?? "Gagal memuat profil."
              : "Tidak dapat terhubung ke server. Coba lagi nanti."
          );
        } else {
          setError("Terjadi kesalahan yang tidak diketahui.");
        }
      } finally {
        if (!cancelled) setLoading(false);
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [navigate]);

  const handlePick = async (kind: "avatar" | "banner", file?: File) => {
    if (!file || mediaBusy) return;
    setMediaError("");
    if (!ACCEPT.split(",").includes(file.type)) {
      setMediaError("Format gambar harus JPEG, PNG, GIF, atau WebP.");
      return;
    }
    if (file.size > LIMIT_MB[kind] * 1024 * 1024) {
      setMediaError(`Ukuran ${kind === "avatar" ? "foto" : "banner"} maksimal ${LIMIT_MB[kind]} MB.`);
      return;
    }
    setMediaBusy(true);
    try {
      setProfile(await uploadProfileImage(kind, file));
    } catch (err) {
      setMediaError(apiErrorMessage(err, "Gagal mengunggah gambar. Coba lagi."));
    } finally {
      setMediaBusy(false);
    }
  };

  const handleRemove = async (kind: "avatar" | "banner") => {
    if (mediaBusy) return;
    setMediaError("");
    setMediaBusy(true);
    try {
      setProfile(await deleteProfileImage(kind));
    } catch (err) {
      setMediaError(apiErrorMessage(err, "Gagal menghapus gambar. Coba lagi."));
    } finally {
      setMediaBusy(false);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/login", { replace: true });
  };

  if (loading) {
    return <div className="profile-state">Loading profile...</div>;
  }

  if (error || !profile) {
    return (
      <div className="profile-state">
        <p className="profile-error" role="alert">
          {error || "Profil tidak ditemukan."}
        </p>
        <button
          className="profile-btn"
          onClick={() => window.location.reload()}
        >
          TRY AGAIN
        </button>
      </div>
    );
  }

  return (
    <div className="profile-page">
      {/* Kolom utama */}
      <section className="profile-main">
        <div
          className="profile-banner"
          style={profile.banner_url ? { backgroundImage: `url("${mediaSrc(profile.banner_url)}")` } : undefined}
        >
          <div className="profile-banner__actions">
            <button type="button" className="media-btn" disabled={mediaBusy} onClick={() => bannerInput.current?.click()}>
              {profile.banner_url ? "Ganti banner" : "Tambah banner"}
            </button>
            {profile.banner_url && (
              <button type="button" className="media-btn" disabled={mediaBusy} onClick={() => handleRemove("banner")}>
                Hapus
              </button>
            )}
          </div>
          {!profile.banner_url && <img src={mascot} alt="" className="profile-banner__mascot" />}
        </div>

        <div className="profile-photo">
          {profile.avatar_url ? (
            <img src={mediaSrc(profile.avatar_url)} alt={`Foto profil ${profile.username}`} className="profile-photo__img" />
          ) : (
            <span className="profile-photo__initial" aria-label={`Inisial ${profile.username}`}>
              {profile.username.charAt(0).toUpperCase()}
            </span>
          )}
          <div className="profile-photo__actions">
            <button type="button" className="media-btn" disabled={mediaBusy} onClick={() => avatarInput.current?.click()}>
              {profile.avatar_url ? "Ganti foto" : "Unggah foto"}
            </button>
            {profile.avatar_url && (
              <button type="button" className="media-btn" disabled={mediaBusy} onClick={() => handleRemove("avatar")}>
                Hapus foto
              </button>
            )}
          </div>
        </div>

        <input
          ref={avatarInput}
          type="file"
          accept={ACCEPT}
          hidden
          onChange={(e) => {
            void handlePick("avatar", e.target.files?.[0]);
            e.target.value = "";
          }}
        />
        <input
          ref={bannerInput}
          type="file"
          accept={ACCEPT}
          hidden
          onChange={(e) => {
            void handlePick("banner", e.target.files?.[0]);
            e.target.value = "";
          }}
        />
        {mediaBusy && <p className="profile-media-note" role="status">Menyimpan gambar...</p>}
        {mediaError && <p className="profile-error profile-media-note" role="alert">{mediaError}</p>}

        <div className="profile-info">
          <h1 className="profile-name">{profile.name}</h1>
          <p className="profile-email">
            @{profile.username} · {profile.email}
          </p>
          <p className="profile-email">
            <b>{profile.followers_count}</b> pengikut · <b>{profile.following_count}</b> mengikuti
          </p>
        </div>

        <hr className="profile-divider" />

        <h2 className="profile-heading">Statistics</h2>

        <div className="stats-grid">
          <div className="stat-card">
            <span className="stat-icon">🔥</span>
            <div>
              <p className="stat-value">{show(stats?.streak)}</p>
              <p className="stat-label">Day streak</p>
            </div>
          </div>

          <div className="stat-card">
            <span className="stat-icon">⚡</span>
            <div>
              <p className="stat-value">{show(profile.xp)}</p>
              <p className="stat-label">Total XP</p>
            </div>
          </div>

          <div className="stat-card">
            <span className="stat-icon">🏅</span>
            <div>
              <p className="stat-value">{show(profile.level)}</p>
              <p className="stat-label">Level</p>
            </div>
          </div>

          <div className="stat-card">
            <span className="stat-icon">❤️</span>
            <div>
              <p className="stat-value">{show(stats?.hearts ?? profile.lives)}</p>
              <p className="stat-label">Hearts</p>
            </div>
          </div>
        </div>

        <h2 className="profile-heading">Achievements</h2>

        <div className="empty-card">
          <p>No achievements yet. Keep learning to unlock them!</p>
        </div>
      </section>

      {/* Kolom kanan */}
      <aside className="profile-side">
        <div className="topbar-stats">
          <span>🔥 {show(stats?.streak)}</span>
          <span>⚡ {show(profile.xp)}</span>
          <span>❤️ {show(stats?.hearts ?? profile.lives)}</span>
        </div>

        <div className="side-card">
          <h3>Account</h3>
          <p className="side-text">{profile.email}</p>

          <button
            className="profile-btn profile-btn-outline"
            onClick={handleLogout}
          >
            LOG OUT
          </button>
        </div>
      </aside>
    </div>
  );
}
