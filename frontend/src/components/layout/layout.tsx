import { useAppContext } from "@/context/app-context";
import { LanguageSelector } from "../language/language";
import { Outlet } from "react-router";

export const Layout = () => {
  const { backgroundImage, backgroundColor, disableLanguageSelector } = useAppContext();

  return (
    <div
      className="relative flex flex-col justify-center items-center min-h-svh"
      style={{
        backgroundColor: backgroundColor || undefined,
        backgroundImage: backgroundImage ? `url(${backgroundImage})` : undefined,
        backgroundSize: "cover",
        backgroundPosition: "center",
      }}
    >
      {!disableLanguageSelector && <LanguageSelector />}
      <Outlet />
    </div>
  );
};
