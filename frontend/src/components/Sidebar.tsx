// frontend/src/components/Sidebar.tsx
import { useEffect, useState, type ReactNode } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import logo from "../assets/logo.png";
import { HomeIcon, CodeTileIcon, ShieldIcon, ChestIcon, StoreIcon, CLangIcon } from "./SidebarIcons";
import { PROFILE_CHANGED_EVENT, getMyProfile } from "../services/profileService";
import { mediaSrc } from "../services/mediaUrl";
import type { Profile } from "../types/profile";
import "./Sidebar.css";

type MenuItem = {
  label: string;
  icon?: ReactNode; // komponen SVG dari SidebarIcons.tsx; kosong untuk menu Profil
  path: string; // halaman tujuan
  // prefix lain yang juga menandai menu ini aktif
  alsoActiveFor?: string[];
  // round = lingkaran avatar profil
  variant?: "round";
};

// Props: semua opsional, jadi <Sidebar /> tetap bisa dipakai tanpa argumen
type SidebarProps = {
  defaultActive?: string; // dipakai bila alamat saat ini tidak cocok dengan menu mana pun
  onSelect?: (label: string) => void; // dipanggil setiap menu diklik
};

// Navigasi utama. Tambah atau hapus menu cukup lewat array ini.
const menuItems: MenuItem[] = [
  { label: "Belajar", icon: <HomeIcon />, path: "/belajar" },
  { label: "Karakter", icon: <CodeTileIcon />, path: "/karakter" },
  { label: "Papan Skor", icon: <ShieldIcon />, path: "/papan-skor", alsoActiveFor: ["/pengguna"] },
  { label: "Misi", icon: <ChestIcon />, path: "/misi" },
  { label: "Toko", icon: <StoreIcon />, path: "/toko" },
  { label: "Profil", path: "/profile", variant: "round" },
];

function matches(pathname: string, prefix: string) {
  return pathname === prefix || pathname.startsWith(prefix + "/");
}

function activeLabel(pathname: string): string | null {
  for (const item of menuItems) {
    if (matches(pathname, item.path) || item.alsoActiveFor?.some((p) => matches(pathname, p))) {
      return item.label;
    }
  }
  return null;
}

// Ikon profil: foto pengguna, atau huruf pertama username bila belum ada foto.
function ProfileAvatar({ profile }: { profile: Profile | null }) {
  const [broken, setBroken] = useState<string>("");
  const src = mediaSrc(profile?.avatar_url);

  if (src && broken !== src) {
    return <img src={src} alt="" className="sidebar__avatar" onError={() => setBroken(src)} />;
  }
  const initial = (profile?.username ?? "").trim().charAt(0).toUpperCase();
  return <>{initial || "?"}</>;
}

export default function Sidebar({ defaultActive = "Belajar", onSelect }: SidebarProps) {
  const navigate = useNavigate();
  const { pathname } = useLocation();
  const active = activeLabel(pathname) ?? defaultActive;

  const [profile, setProfile] = useState<Profile | null>(null);

  useEffect(() => {
    let cancelled = false;
    getMyProfile()
      .then((p) => {
        if (!cancelled) setProfile(p);
      })
      .catch(() => {
        /* ikon cadangan tetap tampil */
      });

    // Foto diganti/dihapus di halaman Profil: perbarui ikon tanpa muat ulang.
    const onChanged = (e: Event) => setProfile((e as CustomEvent<Profile>).detail);
    window.addEventListener(PROFILE_CHANGED_EVENT, onChanged);
    return () => {
      cancelled = true;
      window.removeEventListener(PROFILE_CHANGED_EVENT, onChanged);
    };
  }, []);

  const handleClick = (item: MenuItem) => {
    onSelect?.(item.label);
    navigate(item.path);
  };

  return (
    <aside className="sidebar">
      <div className="sidebar__brand">
        <img src={logo} alt="" className="sidebar__logo" />
        codelingo
      </div>

      <nav aria-label="Menu utama">
        <ul className="sidebar__list">
          {menuItems.map((item) => {
            const isActive = active === item.label;
            return (
              <li key={item.label}>
                <button
                  type="button"
                  className={`sidebar__item${isActive ? " sidebar__item--active" : ""}`}
                  aria-current={isActive ? "page" : undefined}
                  onClick={() => handleClick(item)}
                >
                  <span
                    className={`sidebar__icon${item.variant ? ` sidebar__icon--${item.variant}` : ""}`}
                    aria-hidden="true"
                  >
                    {item.variant === "round" ? <ProfileAvatar profile={profile} /> : item.icon}
                  </span>
                  <span className="sidebar__label">{item.label}</span>
                </button>
              </li>
            );
          })}
        </ul>
      </nav>

      {/* Kartu promo di bawah menu */}
      <div className="sidebar__promo">
        <div className="sidebar__promo-icon" aria-hidden="true">
          <CLangIcon />
        </div>
        <h3>Mau belajar bahasa C?</h3>
        <p>Jadi mudah dengan CodeLingo!</p>
        <button type="button" className="sidebar__promo-btn" onClick={() => navigate("/belajar")}>
          Coba bahasa C
        </button>
      </div>
    </aside>
  );
}
