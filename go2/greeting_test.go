package GO2
import(
	"testing"
	"regexp"
)

func TestHello(t *testing.T) {
	name:= "Cikal Prameswari"
	want := regexp.MustCompile(`\b`+name+`b`)
	msg, err:= Hello("Cikal Prameswari")
	if !want.MatchString(msg) || err != nil{
		t.Errorf(`Hello("Cikal Prameswari") = %q, %v, want match for %#q, nil`, msg,err,want,)
	}
}

func Emp(t *testing.T){
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg,err,)
	}
}