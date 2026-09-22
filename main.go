package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// 1. تقديم الملفات الإستاتيكية (HTML/CSS)
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// 2. اللوجيك الحاي لشريط البحث الحي (Instant Live Search)
	http.HandleFunc("/api/search", handleSearch)

	// 3. اللوجيك الحاي للتنقل بين الأقسام دون إعادة تحميل الصفحة
	http.HandleFunc("/api/view/", handleViewSwitch)

	// 4. اللوجيك الخاص بالملف الشخصي
	http.HandleFunc("/api/profile", handleProfile)

	fmt.Println("🚀 IVcode server active on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// معالج البحث المباشر
func handleSearch(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query == "" {
		fmt.Fprint(w, "")
		return
	}

	w.Header().Set("Content-Type", "text/html")
	
	// محاكاة نتائج بحث ديناميكية بترجع HTML جاهز
	fmt.Fprintf(w, `
		<div style="background: rgba(15, 15, 20, 0.95); border: 1px solid rgba(255,255,255,0.15); border-radius: 8px; padding: 12px; box-shadow: 0 10px 25px rgba(0,0,0,0.8); backdrop-filter: blur(10px);">
			<div style="font-size: 0.85rem; color: #8b949e; margin-bottom: 6px;">Search Results for: <strong style="color: #fff;">%s</strong></div>
			<div style="padding: 6px 0; border-bottom: 1px solid rgba(255,255,255,0.05); font-size: 0.9rem; color: #58a6ff; cursor: pointer;">🔍 gitIV repository: %s-core</div>
			<div style="padding: 6px 0; font-size: 0.9rem; color: #58a6ff; cursor: pointer;">📦 IVstore item: Custom Hardware %s</div>
		</div>
	`, query, query, query)
}

// معالج التنقل الديناميكي بين الأقسام
func handleViewSwitch(w http.ResponseWriter, r *http.Request) {
	section := strings.TrimPrefix(r.URL.Path, "/api/view/")
	w.Header().Set("Content-Type", "text/html")

	switch section {
	case "gitIV":
		fmt.Fprint(w, `
			<div class="card" style="grid-column: 1 / -1;">
				<div class="card-header">
					<h3>gitIV Platform</h3>
					<span class="card-tag">Core Version Control</span>
				</div>
				<p>Git repository ecosystem built for high performance code hosting, automated builds, and deployment management.</p>
			</div>
		`)
	case "IVstore":
		fmt.Fprint(w, `
			<div class="card" style="grid-column: 1 / -1;">
				<div class="card-header">
					<h3>IVstore Marketplace</h3>
					<span class="card-tag">Hardware Marketplace</span>
				</div>
				<p>A dedicated hardware marketplace for selling and acquiring customized hardware components, rigs, and tailored tech equipment exclusively from IVcode.</p>
			</div>
		`)
	case "for-cus":
		fmt.Fprint(w, `
			<div class="card" style="grid-column: 1 / -1;">
				<div class="card-header">
					<h3>for cus Solutions</h3>
					<span class="card-tag">Custom Engineering</span>
				</div>
				<p>Custom software systems designed and engineered specifically for enterprise requirements and client workflows.</p>
			</div>
		`)
	case "IVs-devss":
		fmt.Fprint(w, `
			<div class="card" style="grid-column: 1 / -1;">
				<div class="card-header">
					<h3>IVs devss Hub</h3>
					<span class="card-tag">Developer Ecosystem</span>
				</div>
				<p>Documentation, low-level APIs, SDK downloads, and core development tools for building on top of IVcode infrastructure.</p>
			</div>
		`)
	default:
		// إرجاع الأربعة كروت الأصلية
		fmt.Fprint(w, `
			<div class="card" id="gitIV">
				<div class="card-header">
					<h3>gitIV</h3>
					<span class="card-tag">Platform</span>
				</div>
				<p>The core version-control and source code management platform designed for hosting repositories, managing code releases, and syncing open-source builds.</p>
			</div>
			<div class="card" id="IVstore">
				<div class="card-header">
					<h3>IVstore</h3>
					<span class="card-tag">Marketplace</span>
				</div>
				<p>A dedicated hardware marketplace for selling and acquiring customized hardware components, rigs, and tailored tech equipment exclusively from IVcode.</p>
			</div>
			<div class="card" id="for-cus">
				<div class="card-header">
					<h3>for cus</h3>
					<span class="card-tag">Solutions</span>
				</div>
				<p>Custom tailored software solutions, personalized developer tools, and client-centric workflows engineered to address specific user needs.</p>
			</div>
			<div class="card" id="IVs-devss">
				<div class="card-header">
					<h3>IVs devss</h3>
					<span class="card-tag">Ecosystem</span>
				</div>
				<p>The official developer hub and ecosystem for builders, hosting SDKs, low-level technical documentation, and community tools.</p>
			</div>
		`)
	}
}

// معالج بروفايل الـ CEO
func handleProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<div class="card" style="grid-column: 1 / -1; text-align: center;">
			<div class="card-header" style="justify-content: center; gap: 10px;">
				<h3>CEO Profile</h3>
				<span class="card-tag">Solo Founder</span>
			</div>
			<p style="margin-top: 10px;">Welcome! Working endlessly (solo as the CEO) to satisfy and empower every user across the IVcode ecosystem.</p>
		</div>
	`)
}
