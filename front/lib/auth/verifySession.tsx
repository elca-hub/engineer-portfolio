import { cookieTokenName } from '@/constants/cookieTokenName'
import { redirect } from 'next/navigation'
import { cookies } from 'next/headers'
import 'server-only'

export async function isLogin() {
  return false
  const cookieStore = await cookies()
  const session = cookieStore.get(cookieTokenName)

  if (session) return true
  return false
}

export async function verifySession() {
  const isAlreadyLogin = await isLogin()

  if (!isAlreadyLogin) redirect('./login')
}