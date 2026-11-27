import api from "./api";
import type { PublicProfile, UserList } from "../types/social";

const enc = encodeURIComponent;

export async function getPublicProfile(username: string): Promise<PublicProfile> {
  const res = await api.get<PublicProfile>(`/users/${enc(username)}`);
  return res.data;
}

export async function followUser(username: string): Promise<PublicProfile> {
  const res = await api.post<PublicProfile>(`/users/${enc(username)}/follow`);
  return res.data;
}

export async function unfollowUser(username: string): Promise<PublicProfile> {
  const res = await api.delete<PublicProfile>(`/users/${enc(username)}/follow`);
  return res.data;
}

export async function getFollowList(
  username: string,
  kind: "followers" | "following",
  limit = 50,
): Promise<UserList> {
  const res = await api.get<UserList>(`/users/${enc(username)}/${kind}`, { params: { limit } });
  return res.data;
}
