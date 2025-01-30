import {
  Button,
  Paper,
  PasswordInput,
  TextInput,
  Text,
  Divider,
  Grid,
} from "@mantine/core";
import { useForm, zodResolver } from "@mantine/form";
import { notifications } from "@mantine/notifications";
import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import { z } from "zod";
import { useUserContext } from "../context/user-context";
import { Navigate } from "react-router";
import { Layout } from "../components/layouts/layout";
import { GoogleIcon } from "../icons/google";
import { GithubIcon } from "../icons/github";
import { OAuthIcon } from "../icons/oauth";
import { Logo } from "../components/Logo";

export const LoginPage = () => {
  const queryString = window.location.search;
  const params = new URLSearchParams(queryString);
  const redirectUri = params.get("redirect_uri");

  const { isLoggedIn, configuredProviders } = useUserContext();

  if (isLoggedIn) {
    return <Navigate to="/logout" />;
  }

  const schema = z.object({
    username: z.string(),
    password: z.string(),
  });

  type FormValues = z.infer<typeof schema>;

  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      username: "",
      password: "",
    },
    validate: zodResolver(schema),
  });

  const loginMutation = useMutation({
    mutationFn: (login: FormValues) => {
      return axios.post("/api/login", login);
    },
    onError: () => {
      notifications.show({
        title: "Failed to login",
        message: "Check your username and password",
        color: "red",
      });
    },
    onSuccess: () => {
      notifications.show({
        title: "Logged in",
        message: "Welcome back!",
        color: "green",
      });
      setTimeout(() => {
        if (redirectUri === "null") {
          window.location.replace("/");
        } else {
          window.location.replace(`/continue?redirect_uri=${redirectUri}`);
        }
      }, 500);
    },
  });

  const loginOAuthMutation = useMutation({
    mutationFn: (provider: string) => {
      return axios.get(
        `/api/oauth/url/${provider}?redirect_uri=${redirectUri}`,
      );
    },
    onError: () => {
      notifications.show({
        title: "Internal error",
        message: "Failed to get OAuth URL",
        color: "red",
      });
    },
    onSuccess: (data) => {
      notifications.show({
        title: "Redirecting",
        message: "Redirecting to your OAuth provider",
        color: "blue",
      });
      setTimeout(() => {
        window.location.href = data.data.url;
      }, 500);
    },
  });

  const handleSubmit = (values: FormValues) => {
    loginMutation.mutate(values);
  };

  return (
    <Layout>
      <Logo />
      <Paper shadow="none" p="xl" mt={30} radius="md" withBorder={false}>
        {configuredProviders.length === 0 && (
          <Text size="lg" mb="md" fw={500} ta="center">
            Single Sign-On
          </Text>
        )}
        {configuredProviders.length > 0 && (
          <>
            <Text size="lg" fw={500} ta="center">
              Login via SSO
            </Text>
            <Grid mb="md" mt="md" align="center" justify="center">
              {configuredProviders.includes("google") && (
                <Grid.Col span="content">
                  <Button
                    radius="xl"
                    leftSection={
                      <GoogleIcon style={{ width: 14, height: 14 }} />
                    }
                    variant="default"
                    onClick={() => loginOAuthMutation.mutate("google")}
                    loading={loginOAuthMutation.isLoading}
                  >
                    Google
                  </Button>
                </Grid.Col>
              )}
              {configuredProviders.includes("github") && (
                <Grid.Col span="content">
                  <Button
                    radius="xl"
                    leftSection={
                      <GithubIcon style={{ width: 14, height: 14 }} />
                    }
                    variant="default"
                    onClick={() => loginOAuthMutation.mutate("github")}
                    loading={loginOAuthMutation.isLoading}
                  >
                    Github
                  </Button>
                </Grid.Col>
              )}
              {configuredProviders.includes("generic") && (
                <Grid.Col span="content">
                  <Button
                    radius="xl"
                    leftSection={
                      <OAuthIcon style={{ width: 14, height: 14 }} />
                    }
                    variant="default"
                    onClick={() => loginOAuthMutation.mutate("generic")}
                    loading={loginOAuthMutation.isLoading}
                  >
                    District Email
                  </Button>
                </Grid.Col>
              )}
            </Grid>
            {configuredProviders.includes("username") && (
              <Divider
                label="Or continue with password"
                labelPosition="center"
                my="lg"
              />
            )}
          </>
        )}
        {configuredProviders.includes("username") && (
          <form onSubmit={form.onSubmit(handleSubmit)}>
            <TextInput
              label="Username"
              placeholder="bbaggins@njes.org"
              required
              disabled={loginMutation.isLoading}
              key={form.key("username")}
              {...form.getInputProps("username")}
            />
            <PasswordInput
              label="Password"
              placeholder="BagEnd123!"
              required
              mt="md"
              disabled={loginMutation.isLoading}
              key={form.key("password")}
              {...form.getInputProps("password")}
            />
            <Button
              fullWidth
              mt="xl"
              type="submit"
              loading={loginMutation.isLoading}
            >
              Login
            </Button>
          </form>
        )}
      </Paper>
    </Layout>
  );
};
