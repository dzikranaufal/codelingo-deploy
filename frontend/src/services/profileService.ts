import api from "./api";
import type { Profile } from "../types/profile";

// Dipancarkan setiap profil berubah (foto/banner) agar Sidebar ikut diperbarui.
export const PROFILE_CHANGED_EVENT = "codelingo:profile-changed";

function notifyProfileChanged(profile: Profile) {
  window.dispatchEvent(new CustomEvent<Profile>(PROFILE_CHANGED_EVENT, { detail: profile }));
}

export async function getMyProfile(): Promise<Profile> {
  const res = await api.get<Profile>("/profile");
  return res.data;
}

type MediaKind = "avatar" | "banner";

export async function uploadProfileImage(kind: MediaKind, file: File): Promise<Profile> {
  const form = new FormData();
  form.append("file", file);
  const res = await api.post<Profile>(`/profile/${kind}`, form, {
    // Biarkan browser membuat boundary multipart sendiri.
    headers: { "Content-Type": undefined },
    timeout: 60000,
  });
  notifyProfileChanged(res.data);
  return res.data;
}

export async function deleteProfileImage(kind: MediaKind): Promise<Profile> {
  const res = await api.delete<Profile>(`/profile/${kind}`);
  notifyProfileChanged(res.data);
  return res.data;
}
