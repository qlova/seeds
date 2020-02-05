//Package paymentbox provides a (stripe) secured input that accepts payment details.
package paymentbox

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/user"
	"github.com/qlova/seeds/column"
	"github.com/qlova/seeds/textbox"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/setupintent"
)

//Options for configuring the payment box.
type Options struct {
	StripePublishableKey string
}

//Seed is paymentbox.
type Seed struct {
	column.Seed

	CreditCardDetails seed.Seed
	BillingName       textbox.Seed
}

//New returns a new payment box.
func New(options Options) Seed {
	var PaymentBox = Seed{Seed: column.New()}
	{
		PaymentBox.CreditCardDetails = seed.AddTo(PaymentBox)

		PaymentBox.Require("https://js.stripe.com/v3/")

		PaymentBox.OnReady(func(q script.Ctx) {
			q.Javascript(`%v.stripe = Stripe('%v');`, PaymentBox.Ctx(q).Element(), options.StripePublishableKey)
			q.Javascript(`let elements = %v.stripe.elements();`, PaymentBox.Ctx(q).Element())
			q.Javascript(`%v.card = elements.create("card", {}); %[1]v.card.mount("#%v");`, PaymentBox.Ctx(q).Element(), PaymentBox.CreditCardDetails.ID())
		})

		PaymentBox.BillingName = textbox.AddTo(PaymentBox)
		PaymentBox.BillingName.SetPlaceholder("Billing Name")
		PaymentBox.BillingName.Align().Start()
	}

	return PaymentBox
}

//AddTo parent.
func AddTo(parent seed.Interface, options Options) Seed {
	var NumberBox = New(options)
	parent.Root().Add(NumberBox)
	return NumberBox
}

//Ctx is the script context of the paymentbox.
type Ctx struct {
	q script.Ctx

	script.Seed
	BillingName textbox.Ctx
}

//Ctx returns the paymentbox script context.
func (paymentbox Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{q, paymentbox.Seed.Ctx(q), paymentbox.BillingName.Ctx(q)}
}

//Value returns a reference to the payment.
//For example when using stripe, it returns the PaymentMethod ID.
func (paymentbox Ctx) Value() script.String {
	var q = paymentbox.q

	var secret = q.Go(func(u user.Ctx) string {
		intent, err := setupintent.New(&stripe.SetupIntentParams{
			PaymentMethodTypes: []*string{
				stripe.String("card"),
			},
		})

		if err != nil {
			u.Report(err)
			return ""
		}

		return intent.ClientSecret
	}).Wait().String()

	return q.Value(`await %v.stripe.confirmCardSetup(
		%v,
		{
			payment_method: {
				card: %[1]v.card,
				billing_details: {
					name: %v,
				},
			},
		}
		).then(function(result) {
			if (result.error) {
				console.error(result.error.message);
				return "";
			} else {
				return result.setupIntent.payment_method;
			}
	})`, paymentbox.Element(), secret, paymentbox.BillingName.Value()).String()
}
