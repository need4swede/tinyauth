import { useQuery } from "@tanstack/react-query";
import React, { createContext, useContext } from "react";
import { UserContextSchemaType } from "../schemas/user-context-schema";
import axios from "axios";

// Extend the context type to ensure disableContinue is included
type EnhancedUserContextType = UserContextSchemaType & {
  disableContinue: boolean;
};

const UserContext = createContext<EnhancedUserContextType | null>(null);

export const UserContextProvider = ({
  children,
  autoRedirect = true, // Add prop to control behavior
}: {
  children: React.ReactNode;
  autoRedirect?: boolean;
}) => {
  const {
    data: userContextData,
    isLoading,
    error,
  } = useQuery({
    queryKey: ["isLoggedIn"],
    queryFn: async () => {
      const res = await axios.get("/api/status");
      return res.data;
    },
  });

  if (error && !isLoading) {
    throw error;
  }

  // Enhance the context with disableContinue
  const enhancedContext: EnhancedUserContextType | null = userContextData
    ? {
        ...userContextData,
        disableContinue: autoRedirect, // Set based on prop
      }
    : null;

  return (
    <UserContext.Provider value={enhancedContext}>
      {children}
    </UserContext.Provider>
  );
};

export const useUserContext = () => {
  const context = useContext(UserContext);
  if (context === null) {
    throw new Error("useUserContext must be used within a UserContextProvider");
  }
  return context;
};