export interface PublicProfile {
  username: string;
  name: string;
  level: number;
  xp: number;
  avatar_url: string;
  banner_url: string;
  joined_at: string;
  followers_count: number;
  following_count: number;
  is_following: boolean;
  follows_you: boolean;
  is_me: boolean;
}

export interface UserCard {
  username: string;
  name: string;
  xp: number;
  avatar_url: string;
  is_following: boolean;
  is_me: boolean;
}

export interface UserList {
  users: UserCard[];
  total: number;
  limit: number;
  offset: number;
}
