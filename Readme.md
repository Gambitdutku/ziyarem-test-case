# Ziyarem Test Case – Smart City Sensor Data Platform

Bu proje, **Akıllı Şehir Sensör Platformu** test case’ine uygun bir Go backend uygulamasıdır.  
DDD + Hexagonal mimari yaklaşımı ile tasarlanmıştır.

---

##  Özellikler
- Sensor interface + 3 farklı implementasyon (Temperature, Humidity, AirQuality)
- Redis cache kontrolü + DB fallback
- Error handling (işsel hata vs sistem hatası)
- Retry mekanizması + Circuit Breaker (3 hata → 30s open)
- Loglama (`info`, `warn`, `error` seviyeleri, logrus)
- Auto migration ile DB tabloları otomatik oluşturulur

---

##  Proje Yapısı
```
ziyaremtestcase/
├── cmd/main.go               # giriş noktası
├── domain/                   # domain modelleri ve interface
│   ├── sensor.go
│   ├── device.go
│   ├── type.go
│   └── api.go
├── sensors/                  # farklı sensör implementasyonları
│   ├── temp_sensor.go
│   ├── humidity_sensor.go
│   └── airquality_sensor.go
├── application/              # iş mantığı devreleri
│   ├── service.go
│   └── circuitbreaker.go
├── infrastructure/           # DB ve cache adaptörrleri
│   ├── db.go
│   ├── cache.go
│   └── repository.go
├── go.mod / go.sum
├── README.md
└── mockapi.py                #simule etmek için 
```

---

##  Gereksinimler
- Go 1.20+
- MySQL 8+
- Redis 7+
- Python 3 (mock sensör API için)

---

##  Kurulum

### MySQL
```bash
sudo apt install mysql-server -y
mysql -u root -p
mysql -u username -p < zirayem_dump.sql #gerekli tabloları programı oluşturudğumuzda eklemekte fakat gerekli verileri ekleyecek bir fonksiyon bulunmadığı için dump kullanılması gerekiyor
```

### Redis
```bash
sudo apt install redis-server -y
redis-cli ping
# PONG
```

### Go bağımlılıkları
```bash
go mod tidy
```

---

##  Çalıştırma
```bash
go run ./cmd/main.go
```

Log çıktısı:
```
INFO[0000] Cache hit: temperature:sensor-123
INFO[0000] SensorData okundu ve kaydedildi: {...}
```

---


```bash
pip install flask
python3 mock_api.py
```

- `http://localhost:8081/temp/temp-001`
- `http://localhost:8081/hum/hum-001`
- `http://localhost:8081/air/air-001`

adresleri üzerinden sahte sensör verileri döner.

---

##  Sonuç
Bu proje, verilen test case’in tüm gereksinimlerini karşılamaktadır:
- DDD/Hexagonal mimari
- Sensor interface + implementasyonlar
- Redis + DB entegrasyonu
- Error handling & resilience
- Circuit breaker
- Loglama  