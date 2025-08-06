/**
 * ハイフンを年月日に変換する
 * @param date 日付('2024-11-16')
 *
 * @returns 日付('2024年11月16日')
 */
export default function formatDate(date: string) {
	return date.replace(/-/g, '年').replace(/-/g, '月') + '日'
}
