import { Loader2 } from "lucide-react";
import { Button } from "./button";
import React from "react";
import { twMerge } from "tailwind-merge";

interface Props extends React.ComponentProps<typeof Button> {
  title: string;
  icon: React.ReactNode;
  onClick?: () => void;
  loading?: boolean;
  textColor?: string;
  backgroundColor?: string;
  hideIcon?: boolean;
}

export const OAuthButton = (props: Props) => {
  const { title, icon, onClick, loading, className, textColor, backgroundColor, hideIcon, ...rest } = props;

  const customStyle = {
    color: textColor || undefined,
    backgroundColor: backgroundColor || undefined,
    borderColor: backgroundColor || undefined,
  };

  return (
    <Button
      onClick={onClick}
      className={twMerge("rounded-md", className)}
      variant="outline"
      style={customStyle}
      {...rest}
    >
      {loading ? (
        <Loader2 className="animate-spin" />
      ) : (
        <>
          {!hideIcon && icon}
          {title}
        </>
      )}
    </Button>
  );
};
