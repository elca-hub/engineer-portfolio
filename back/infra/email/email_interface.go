//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package email

type Email interface {
	SendEmail(to string, subject string, vars interface{}, files ...string) error
}
