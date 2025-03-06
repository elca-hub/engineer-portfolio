import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  experimental: {
    optimizePackageImports: ["@/components/ui"],
  },
};

export default nextConfig;
