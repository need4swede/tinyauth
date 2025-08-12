import { useEffect } from "react";
import { useAppContext } from "@/context/app-context";

export const DynamicHead = () => {
  const { title, favicon, backgroundColor } = useAppContext();

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

  useEffect(() => {
    // Update background color
    if (backgroundColor) {
      // Set the CSS custom property --background on the document root
      document.documentElement.style.setProperty('--background', backgroundColor);
    } else {
      // Remove the custom property if no background color is set
      document.documentElement.style.removeProperty('--background');
    }
  }, [backgroundColor]);

  // This component doesn't render anything visible
  return null;
};
