import { useEffect } from "react";
import { useAppContext } from "@/context/app-context";

export const DynamicHead = () => {
  const { title, favicon } = useAppContext();

  useEffect(() => {
    // Update document title
    if (title) {
      document.title = title;
    }
  }, [title]);

  useEffect(() => {
    // Update favicon
    if (favicon) {
      // Find existing favicon links and update them
      const existingFavicons = document.querySelectorAll('link[rel*="icon"]');
      existingFavicons.forEach((link) => {
        (link as HTMLLinkElement).href = favicon;
      });

      // If no favicon links exist, create a new one
      if (existingFavicons.length === 0) {
        const link = document.createElement('link');
        link.rel = 'icon';
        link.href = favicon;
        document.head.appendChild(link);
      }
    }
  }, [favicon]);

  // This component doesn't render anything visible
  return null;
};
