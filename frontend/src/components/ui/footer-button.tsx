import { useAppContext } from "@/context/app-context";
import { Button } from "./button";

export const FooterButton = () => {
  const {
    footerButtonText,
    footerButtonUrl,
    footerButtonNewTab,
    footerButtonTextColor,
    footerButtonBackgroundColor
  } = useAppContext();

  // Don't render if text or URL is missing
  if (!footerButtonText || !footerButtonUrl) {
    return null;
  }

  const handleClick = () => {
    if (footerButtonNewTab) {
      window.open(footerButtonUrl, '_blank', 'noopener,noreferrer');
    } else {
      window.location.href = footerButtonUrl;
    }
  };

  return (
    <div className="absolute bottom-4 left-1/2 transform -translate-x-1/2">
      <Button
        type="button"
        variant="outline"
        onClick={handleClick}
        style={{
          color: footerButtonTextColor || undefined,
          backgroundColor: footerButtonBackgroundColor || undefined,
          borderColor: footerButtonBackgroundColor || undefined,
        }}
        className="text-sm px-4 py-2"
      >
        {footerButtonText}
      </Button>
    </div>
  );
};
