export interface Profile {
  id: number;
  username: string;
  name: string;
  email: string;
  age: number;
  role: string;
  xp: number;
  level: number;
  lives: number;
  avatar_url: string; // kosong = belum ada foto
  banner_url: string; // kosong = belum ada banner
  followers_count: number;
  following_count: number;
}
