package verify

import (
	"errors"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var interactive bool
var username string
var password string
var docker bool
var user string

var VerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a user is set up correctly",
	Long:  `Verify a user is set up correctly meaning that it has a correct username and password.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Level(zerolog.InfoLevel)

		if interactive {
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("User (username:hash)").Value(&user).Validate((func(s string) error {
						if s == "" {
							return errors.New("user cannot be empty")
						}
						return nil
					})),
					huh.NewInput().Title("Username").Value(&username).Validate((func(s string) error {
						if s == "" {
							return errors.New("username cannot be empty")
						}
						return nil
					})),
					huh.NewInput().Title("Password").Value(&password).Validate((func(s string) error {
						if s == "" {
							return errors.New("password cannot be empty")
						}
						return nil
					})),
					huh.NewSelect[bool]().Title("Is the user formatted for docker?").Options(huh.NewOption("Yes", true), huh.NewOption("No", false)).Value(&docker),
				),
			)

			var baseTheme *huh.Theme = huh.ThemeBase()

			formErr := form.WithTheme(baseTheme).Run()

			if formErr != nil {
				log.Fatal().Err(formErr).Msg("Form failed")
			}
		}

		if username == "" || password == "" || user == "" {
			log.Fatal().Msg("Username, password and user cannot be empty")
		}

		log.Info().Str("user", user).Str("username", username).Str("password", password).Bool("docker", docker).Msg("Verifying user")

		userSplit := strings.Split(user, ":")

		if userSplit[1] == "" {
			log.Fatal().Msg("User is not formatted correctly")
		}

		if docker {
			userSplit[1] = strings.ReplaceAll(userSplit[1], "$$", "$")
		}

		verifyErr := bcrypt.CompareHashAndPassword([]byte(userSplit[1]), []byte(password))

		if verifyErr != nil || username != userSplit[0] {
			log.Fatal().Msg("Username or password incorrect")
		} else {
			log.Info().Msg("Verification successful")
		}
	},
}

func init() {
	VerifyCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Create a user interactively")
	VerifyCmd.Flags().BoolVar(&docker, "docker", false, "Is the user formatted for docker?")
	VerifyCmd.Flags().StringVar(&username, "username", "", "Username")
	VerifyCmd.Flags().StringVar(&password, "password", "", "Password")
	VerifyCmd.Flags().StringVar(&user, "user", "", "Hash (username:hash combination)")
}
