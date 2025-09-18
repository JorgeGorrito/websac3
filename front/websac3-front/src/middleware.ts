import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

// Define protected routes and their required roles
const protectedRoutes = {
  '/admin': 'admin',
  '/director': ['guest', 'program lead'], 
  '/experto': 'cybersecurity_auditor',
}

// Define public routes that don't require authentication
const publicRoutes = [
  '/login',
  '/access-request',
  '/',
]

export function middleware(request: NextRequest) {
  // TEMPORARILY DISABLED - Let client-side authentication handle everything
  // This prevents server-side redirects that conflict with client-side auth
  return NextResponse.next()
  
  /* ORIGINAL MIDDLEWARE CODE - COMMENTED OUT
  const { pathname } = request.nextUrl
  
  // Check if the route is public
  const isPublicRoute = publicRoutes.some(route => 
    pathname === route || pathname.startsWith(route + '/')
  )
  
  if (isPublicRoute) {
    return NextResponse.next()
  }
  
  // Check if the route is protected
  const protectedRoute = Object.keys(protectedRoutes).find(route => 
    pathname.startsWith(route)
  )
  
  if (protectedRoute) {
    // Get the access token from cookies or headers
    const accessToken = request.cookies.get('access_token')?.value || 
                       request.headers.get('authorization')?.replace('Bearer ', '')
    
    if (!accessToken) {
      // Redirect to login if no token
      const loginUrl = new URL('/login', request.url)
      return NextResponse.redirect(loginUrl)
    }
    
    // Note: In a real application, you would verify the JWT token here
    // For now, we'll just check if it exists
    // You could also decode the JWT and check the role/permissions
  }
  
  return NextResponse.next()
  */
}

export const config = {
  matcher: [
    /*
     * Match all request paths except for the ones starting with:
     * - api (API routes)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico (favicon file)
     */
    '/((?!api|_next/static|_next/image|favicon.ico).*)',
  ],
}
