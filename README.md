# Go Web Sayfası Kaydedici (HTML, Screenshot ve Linkler)

Bu proje, **Go** dili ile yazılmış basit bir komut satırı aracıdır. Verilen bir web sitesine giderek:

* Sayfanın **HTML içeriğini** kaydeder
* Sayfanın **tam ekran görüntüsünü (screenshot)** alır
* Sayfa içindeki **mutlak (http/https) linkleri** çıkarıp `.txt` dosyasına yazar

Proje özellikle **Go öğrenen yeni başlayanlar** için sade, anlaşılır ve öğretici olacak şekilde tasarlanmıştır.

---

## 🚀 Özellikler

* 🌐 Verilen URL’ye otomatik olarak gider
* 📄 HTML içeriğini dosyaya kaydeder
* 🖼️ Tam sayfa screenshot alır
* 🔗 Sayfadaki linkleri listeler ve dosyaya yazar
* 🧼 URL’den **dosya sistemiyle uyumlu** dosya isimleri üretir

---

## 🛠️ Kullanılan Teknolojiler

* **Go** (Golang)
* **chromedp** – Headless Chrome kontrolü

---

## 📦 Kurulum

### 1️⃣ Go yüklü mü kontrol et

```bash
go version
```

### 2️⃣ Projeyi klonla

```bash
git clone <repo-url>
cd proje-klasoru
```

### 3️⃣ Gerekli paketi yükle

```bash
go get github.com/chromedp/chromedp
```

---

## ▶️ Kullanım

Program `-url` parametresi ile çalışır.

```bash
go run main.go -url https://www.google.com
```

---

## 📁 Üretilen Dosyalar

Örnek URL:

```
https://www.google.com
```

Oluşan dosyalar:

```
google_com.txt        -> HTML içeriği
google_com.png        -> Screenshot
google_com_links.txt  -> Sayfada bulunan linkler
```

---

## 🧠 Kod Mantığı (Özet)

1. Komut satırından URL alınır
2. `chromedp` ile siteye gidilir
3. HTML ve screenshot alınır
4. URL güvenli dosya adına dönüştürülür
5. Veriler dosyalara yazılır
6. Regex ile linkler çıkarılır

---

## 🔗 Link Çıkarma

Sadece **mutlak linkler** (`http` / `https`) alınır:

```regex
href=["'](http[^"']+)["']
```

Relative linkler (`/about`, `/contact` vb.) dahil edilmez.
