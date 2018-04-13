package actions

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

var r *render.Engine
var assetsBox = packr.NewBox("../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			// uncomment for non-Bootstrap form helpers:
			// "form":     plush.FormHelper,
			// "form_for": plush.FormForHelper,

			// Not specifically proud of this part but I couldn't in timely manner figure out
			// how to nicely have buffalo context here
			"genderOptions": func(maleTranslation string, femaleTranslation string) map[string]string {
				genders := map[string]string{
					maleTranslation:   "male",
					femaleTranslation: "female",
				}
				return genders
			},
		},
	})
}
