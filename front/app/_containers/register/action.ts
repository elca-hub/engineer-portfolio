"use server";

import { apiPrefix } from "@/constants/api";

type ErrorResponse = {
  status: number;
  errors?: string[];
}

export const registerApi = async (email: string, password: string, birthday: string, name: string): Promise<ErrorResponse> => {
  const res = await fetch(`${apiPrefix}/signup`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email,
      password,
      birthday,
      name,
    }),
    credentials: "include",
    cache: "no-cache",
  });
  
  if (res.ok) {
    return { status: res.status };
  } else {
    const error = await res.json();

    return {
      status: res.status,
      errors: error.errors,
    }
  }
}
