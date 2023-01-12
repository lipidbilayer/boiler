module.exports = {
    output: 'standalone',
    async rewrites() {
        return [
          {
            source: '/api/:path*',
            destination: process.env.NEXT_BACKEND_API_URL+'/api/:path*',
          },
          {
            source: '/auth/:path*',
            destination: process.env.NEXT_BACKEND_API_URL+'/auth/:path*',
          }
        ]
      },
};
