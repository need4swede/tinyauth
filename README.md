<div align="center">
    <img alt="Tinyauth" title="Tinyauth" width="96" src="assets/logo-rounded.png">
    <h1>Tinyauth</h1>
    <p>The easiest way to secure your apps with a login screen.</p>
</div>

<div align="center">
    <img alt="License" src="https://img.shields.io/github/license/steveiliop56/tinyauth">
    <img alt="Release" src="https://img.shields.io/github/v/release/steveiliop56/tinyauth">
    <img alt="Issues" src="https://img.shields.io/github/issues/steveiliop56/tinyauth">
    <img alt="Tinyauth CI" src="https://github.com/steveiliop56/tinyauth/actions/workflows/ci.yml/badge.svg">
    <a title="Crowdin" target="_blank" href="https://crowdin.com/project/tinyauth"><img src="https://badges.crowdin.net/tinyauth/localized.svg"></a>
</div>

<br />

> [!NOTE]
> This is a fork of the original [Tinyauth](https://github.com/steveiliop56/tinyauth) with additional features including **Microsoft OAuth support** and **extensive UI customization options**.

Tinyauth is a simple authentication middleware that adds a simple login screen or OAuth with Google, Github, **Microsoft** and any provider to all of your docker apps. It supports all the popular proxies like Traefik, Nginx and Caddy.

![Screenshot](assets/screenshot.png)

> [!WARNING]
> Tinyauth is in active development and configuration may change often. Please make sure to carefully read the release notes before updating.

## Getting Started

You can easily get started with Tinyauth by following the guide in the [documentation](https://tinyauth.app/docs/getting-started.html). There is also an available [docker compose](./docker-compose.example.yml) file that has Traefik, Whoami and Tinyauth to demonstrate its capabilities.

## New Features in This Fork

This fork extends the original Tinyauth with several powerful new features:

### 🔐 Microsoft OAuth Support
Full Microsoft OAuth integration with configurable endpoints and scopes:
```env
MICROSOFT_CLIENT_ID=your_client_id
MICROSOFT_CLIENT_SECRET=your_client_secret
MICROSOFT_AUTH_URL=https://login.microsoftonline.com/common/oauth2/v2.0/authorize
MICROSOFT_TOKEN_URL=https://login.microsoftonline.com/common/oauth2/v2.0/token
MICROSOFT_USER_URL=https://graph.microsoft.com/v1.0/me
MICROSOFT_SCOPES=https://graph.microsoft.com/User.Read
```

### 🎨 Extensive UI Customization
Complete control over the login interface appearance:

**Basic Branding:**
- Custom app titles and login messages
- Logo and favicon customization
- Background images and colors
- Custom form field labels and placeholders

**Advanced Styling:**
- Individual color customization for all UI elements
- Font sizes and layout options
- Border and language selector toggles
- Custom button styling for all OAuth providers

**Example customization:**
```env
APP_TITLE=My Company Portal
LOGIN_TITLE=Welcome Back
LOGIN_SUBTITLE=Sign in to access your applications
LOGO=https://mycompany.com/logo.png
BACKGROUND_COLOR=#f8f9fa
LOGIN_CARD_COLOR=#ffffff
LOGIN_BUTTON_BACKGROUND_COLOR=#007bff
```

### 🌐 Domain-Specific Branding
Apply different branding based on the accessing domain using the `BRANDING_{DOMAIN}_{SETTING}` pattern:

```env
# Corporate portal styling for app.company.com
BRANDING_APP_COMPANY_COM_APP_TITLE=Corporate Portal
BRANDING_APP_COMPANY_COM_LOGIN_TITLE=Employee Access
BRANDING_APP_COMPANY_COM_BACKGROUND_COLOR=#0066cc

# Admin interface styling for admin.company.com
BRANDING_ADMIN_COMPANY_COM_APP_TITLE=Admin Dashboard
BRANDING_ADMIN_COMPANY_COM_LOGIN_TITLE=Administrator Access
BRANDING_ADMIN_COMPANY_COM_BACKGROUND_COLOR=#1a1a1a
```

### 📝 Complete Configuration Reference
See the [.env.example](.env.example) file for a comprehensive list of all available configuration options, organized into logical sections with detailed examples.

## Demo

If you are still not sure if Tinyauth suits your needs you can try out the [demo](https://demo.tinyauth.app). The default username is `user` and the default password is `password`.

## Documentation

You can find documentation and guides on all of the available configuration of Tinyauth in the [website](https://tinyauth.app).

## Discord

Tinyauth has a [discord](https://discord.gg/eHzVaCzRRd) server. Feel free to hop in to chat about self-hosting, homelabs and of course Tinyauth. See you there!

## Contributing

All contributions to the codebase are welcome! If you have any free time feel free to pick up an [issue](https://github.com/steveiliop56/tinyauth/issues) or add your own missing features. Make sure to check out the [contributing guide](./CONTRIBUTING.md) for instructions on how to get the development server up and running.

## Localization

If you would like to help translate Tinyauth into more languages, visit the [Crowdin](https://crowdin.com/project/tinyauth) page.

## License

Tinyauth is licensed under the GNU General Public License v3.0. TL;DR — You may copy, distribute and modify the software as long as you track changes/dates in source files. Any modifications to or software including (via compiler) GPL-licensed code must also be made available under the GPL along with build & install instructions. For more information about the license check the [license](./LICENSE) file.

## Sponsors

A big thank you to the following people for providing me with more coffee:

<!-- sponsors --><a href="https://github.com/erwinkramer"><img src="https:&#x2F;&#x2F;github.com&#x2F;erwinkramer.png" width="64px" alt="User avatar: erwinkramer" /></a>&nbsp;&nbsp;<a href="https://github.com/nicotsx"><img src="https:&#x2F;&#x2F;github.com&#x2F;nicotsx.png" width="64px" alt="User avatar: nicotsx" /></a>&nbsp;&nbsp;<a href="https://github.com/SimpleHomelab"><img src="https:&#x2F;&#x2F;github.com&#x2F;SimpleHomelab.png" width="64px" alt="User avatar: SimpleHomelab" /></a>&nbsp;&nbsp;<a href="https://github.com/jmadden91"><img src="https:&#x2F;&#x2F;github.com&#x2F;jmadden91.png" width="64px" alt="User avatar: jmadden91" /></a>&nbsp;&nbsp;<a href="https://github.com/tribor"><img src="https:&#x2F;&#x2F;github.com&#x2F;tribor.png" width="64px" alt="User avatar: tribor" /></a>&nbsp;&nbsp;<a href="https://github.com/eliasbenb"><img src="https:&#x2F;&#x2F;github.com&#x2F;eliasbenb.png" width="64px" alt="User avatar: eliasbenb" /></a>&nbsp;&nbsp;<a href="https://github.com/afunworm"><img src="https:&#x2F;&#x2F;github.com&#x2F;afunworm.png" width="64px" alt="User avatar: afunworm" /></a>&nbsp;&nbsp;<!-- sponsors -->

## Acknowledgements

- **Freepik** for providing the police hat and badge.
- **Renee French** for the original gopher logo.
- **Coderabbit AI** for providing free AI code reviews.
- **Syrhu** for providing the background image of the app.

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=steveiliop56/tinyauth&type=Date)](https://www.star-history.com/#steveiliop56/tinyauth&Date)
