import api from "./api";

// URL gambar dari backend berbentuk "/api/v1/media/..." (relatif ke server API).
// Ubah menjadi alamat penuh memakai origin dari baseURL axios.
export function mediaSrc(path?: string | null): string {
  if (!path) return "";
  if (/^https?:\/\//i.test(path)) return path;
  try {
    return new URL(path, api.defaults.baseURL).toString();
  } catch {
    return "";
  }
}
