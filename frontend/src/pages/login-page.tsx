import { LoginForm } from "@/components/auth/login-form";
import { GenericIcon } from "@/components/icons/generic";
import { GithubIcon } from "@/components/icons/github";
import { GoogleIcon } from "@/components/icons/google";
import { MicrosoftIcon } from "@/components/icons/microsoft";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from "@/components/ui/card";
import { OAuthButton } from "@/components/ui/oauth-button";
import { SeperatorWithChildren } from "@/components/ui/separator";
import { useAppContext } from "@/context/app-context";
import { useUserContext } from "@/context/user-context";
import { useIsMounted } from "@/lib/hooks/use-is-mounted";
import { LoginSchema } from "@/schemas/login-schema";
import { useMutation } from "@tanstack/react-query";
import axios, { AxiosError } from "axios";
import { useEffect } from "react";
import { useTranslation } from "react-i18next";
import { Navigate, useLocation } from "react-router";
import { toast } from "sonner";

export const LoginPage = () => {
  const { isLoggedIn } = useUserContext();

  if (isLoggedIn) {
    return <Navigate to="/logout" />;
  }

  const {
    configuredProviders,
    title,
    oauthAutoRedirect,
    genericName,
    loginTitle,
    loginSubtitle,
    logo,
    logoSize,
    loginTitleSize,
    loginSubtitleSize,
    disableBorder,
    loginCardColor,
    loginTitleColor,
    loginSubtitleColor
  } = useAppContext();
  const { search } = useLocation();
  const { t } = useTranslation();
  const isMounted = useIsMounted();

  const searchParams = new URLSearchParams(search);
  const redirectUri = searchParams.get("redirect_uri");

  const oauthConfigured =
    configuredProviders.filter((provider) => provider !== "username").length >
    0;
  const userAuthConfigured = configuredProviders.includes("username");

  const oauthMutation = useMutation({
    mutationFn: (provider: string) =>
      axios.get(
        `/api/oauth/url/${provider}?redirect_uri=${encodeURIComponent(redirectUri ?? "")}`,
      ),
    mutationKey: ["oauth"],
    onSuccess: (data) => {
      toast.info(t("loginOauthSuccessTitle"), {
        description: t("loginOauthSuccessSubtitle"),
      });

      setTimeout(() => {
        window.location.href = data.data.url;
      }, 500);
    },
    onError: () => {
      toast.error(t("loginOauthFailTitle"), {
        description: t("loginOauthFailSubtitle"),
      });
    },
  });

  const loginMutation = useMutation({
    mutationFn: (values: LoginSchema) => axios.post("/api/login", values),
    mutationKey: ["login"],
    onSuccess: (data) => {
      if (data.data.totpPending) {
        window.location.replace(
          `/totp?redirect_uri=${encodeURIComponent(redirectUri ?? "")}`,
        );
        return;
      }

      toast.success(t("loginSuccessTitle"), {
        description: t("loginSuccessSubtitle"),
      });

      setTimeout(() => {
        window.location.replace(
          `/continue?redirect_uri=${encodeURIComponent(redirectUri ?? "")}`,
        );
      }, 500);
    },
    onError: (error: AxiosError) => {
      toast.error(t("loginFailTitle"), {
        description:
          error.response?.status === 429
            ? t("loginFailRateLimit")
            : t("loginFailSubtitle"),
      });
    },
  });

  useEffect(() => {
    if (isMounted()) {
      if (
        oauthConfigured &&
        configuredProviders.includes(oauthAutoRedirect) &&
        redirectUri
      ) {
        oauthMutation.mutate(oauthAutoRedirect);
      }
    }
  }, []);

  return (
    <Card
      className={`min-w-xs sm:min-w-sm ${disableBorder ? 'border-0' : ''}`}
      style={{
        backgroundColor: loginCardColor || undefined,
      }}
    >
      <CardHeader>
        {logo && (
          <div className="flex justify-center mb-4">
            {(() => {
              const sizeMap = {
                xs: "h-8",
                sm: "h-12",
                md: "h-16",
                lg: "h-20",
                xl: "h-24",
                "2xl": "h-32",
                "3xl": "h-40",
                "4xl": "h-48",
              };
              const size = (logoSize && logoSize in sizeMap ? logoSize : "md") as keyof typeof sizeMap;
              return (
                <img
                  src={logo}
                  alt="Logo"
                  className={`${sizeMap[size] || "h-16"} w-auto object-contain`}
                  onError={(e) => {
                    e.currentTarget.style.display = 'none';
                  }}
                />
              );
            })()}
          </div>
        )}
        {(() => {
          const sizeMap = {
            sm: "text-sm",
            md: "text-md",
            lg: "text-lg",
            xl: "text-xl",
            "2xl": "text-2xl",
            "3xl": "text-3xl",
            "4xl": "text-4xl",
          };
          const titleSize = (loginTitleSize && loginTitleSize in sizeMap ? loginTitleSize : "3xl") as keyof typeof sizeMap;
          return (
            <CardTitle
              className={`text-center ${sizeMap[titleSize] || "text-3xl"}`}
              style={{
                color: loginTitleColor || undefined,
              }}
            >
              {loginTitle || title}
            </CardTitle>
          );
        })()}
        {configuredProviders.length > 0 &&
          (() => {
            const sizeMap = {
              sm: "text-sm",
              md: "text-md",
              lg: "text-lg",
              xl: "text-xl",
              "2xl": "text-2xl",
              "3xl": "text-3xl",
              "4xl": "text-4xl",
            };
            const subtitleSize = (loginSubtitleSize && loginSubtitleSize in sizeMap ? loginSubtitleSize : "lg") as keyof typeof sizeMap;
            return (
              <CardDescription
                className={`text-center ${sizeMap[subtitleSize] || "text-lg"}`}
                style={{
                  color: loginSubtitleColor || undefined,
                }}
              >
                {loginSubtitle || (oauthConfigured ? t("loginTitle") : t("loginTitleSimple"))}
              </CardDescription>
            );
          })()}
      </CardHeader>
      <CardContent className="flex flex-col gap-4">
        {oauthConfigured && (
          <div className="flex flex-col gap-2 items-center justify-center">
            {configuredProviders.includes("google") && (
              <OAuthButton
                title="Google"
                icon={<GoogleIcon />}
                className="w-full"
                onClick={() => oauthMutation.mutate("google")}
                loading={oauthMutation.isPending && oauthMutation.variables === "google"}
                disabled={oauthMutation.isPending || loginMutation.isPending}
              />
            )}
            {configuredProviders.includes("microsoft") && (
              <OAuthButton
                title="Microsoft"
                icon={<MicrosoftIcon />}
                className="w-full"
                onClick={() => oauthMutation.mutate("microsoft")}
                loading={oauthMutation.isPending && oauthMutation.variables === "microsoft"}
                disabled={oauthMutation.isPending || loginMutation.isPending}
              />
            )}
            {configuredProviders.includes("github") && (
              <OAuthButton
                title="Github"
                icon={<GithubIcon />}
                className="w-full"
                onClick={() => oauthMutation.mutate("github")}
                loading={oauthMutation.isPending && oauthMutation.variables === "github"}
                disabled={oauthMutation.isPending || loginMutation.isPending}
              />
            )}
            {configuredProviders.includes("generic") && (
              <OAuthButton
                title={genericName}
                icon={<GenericIcon />}
                className="w-full"
                onClick={() => oauthMutation.mutate("generic")}
                loading={oauthMutation.isPending && oauthMutation.variables === "generic"}
                disabled={oauthMutation.isPending || loginMutation.isPending}
              />
            )}
          </div>
        )}
        {userAuthConfigured && oauthConfigured && (
          <SeperatorWithChildren>{t("loginDivider")}</SeperatorWithChildren>
        )}
        {userAuthConfigured && (
          <LoginForm
            onSubmit={(values) => loginMutation.mutate(values)}
            loading={loginMutation.isPending || oauthMutation.isPending}
          />
        )}
        {configuredProviders.length == 0 && (
          <p className="text-center text-red-600 max-w-sm">
            {t("failedToFetchProvidersTitle")}
          </p>
        )}
      </CardContent>
    </Card>
  );
};
