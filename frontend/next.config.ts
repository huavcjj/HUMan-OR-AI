import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'lh3.googleusercontent.com',
      },
    ],
  },
  webpack: (config) => {
    config.module.rules.push({
      test: /lucide-react/,
      exclude: /node_modules/,
    });
    return config;
  },
};

export default nextConfig;
