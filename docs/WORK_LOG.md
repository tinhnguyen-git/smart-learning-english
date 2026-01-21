# WORK LOG & STATUS

## CURRENT STATUS

### 📱 Mobile (Flutter)
- **Phase:** Auth Integration
- **Last Action:** (Restoring context) Implemented Auth Bloc & Repository.
- **Next Up:** Login UI & Integration with Backend.

### 🌐 Web (Angular Admin)
- **Phase:** Core Features Integration
- **Last Action:** Integrated Web Auth (Login/Register), Story Scraper, and User Management.
- **Next Up:** Payment Integration or Story Approval Flow.

### ⚙️ Backend (Go)
- **Phase:** API Support
- **Last Action:** Enabled CORS, User Upgrade Endpoint.
- **Next Up:** Payment Service Enhancements.

## FEATURE TRACKER
| Feature | Mobile | Web | Backend | Status |
| :--- | :--- | :--- | :--- | :--- |
| Auth (JWT) | ⬜ | ✅ | ✅ | Done |
| Story Scraper | N/A | ✅ | ✅ | Done |
| Premium Upgrade | ⬜ | ✅ | ✅ | Done |

## KNOWN ISSUES
- **Scraped stories not appearing in Story Approval**: Stories imported via Scraper are saved to DB but don't show up in the Story Approval UI (needs API integration).
- **No English Learning conversion**: Scraped content is not yet converted to English learning materials (vocabulary extraction, difficulty analysis, etc.).

## LOG HISTORY
- **2026-01-21**: Integrated Web Frontend with Backend.
    - Added `app.config.ts` with `provideHttpClient`.
    - Implemented `AuthService`, `StoryService`, `UserService`.
    - Added Login, Register, Profile pages.
    - Fixed DB_HOST to use container name `postgres` after joining devcontainer to `docker_default` network.
    - Enabled CORS on Backend.
    - Upgraded Scraper to use `rod` headless browser (bypasses Cloudflare/JS-rendered pages).
    - Fixed Logout button in sidebar.