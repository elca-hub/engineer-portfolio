"use server";

import { isLogin } from '@/lib/auth/verifySession'
import { redirect } from 'next/navigation';
import UserLoginPresentation from './presentation';

export default async function UserLoginContainer() {
  const isAlreadyLogin = await isLogin();

  if (isAlreadyLogin) {
    redirect('./user')
  }

  return (
    <UserLoginPresentation></UserLoginPresentation>
  )
}
