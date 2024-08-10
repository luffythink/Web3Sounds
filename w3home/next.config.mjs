/** @type {import('next').NextConfig} */
const nextConfig = {
  eslint: {
    ignoreDuringBuilds: true
  },
  webpack: (config) => {
    config.externals.push('pino-pretty');

    return config;
  }
};

export default nextConfig;
