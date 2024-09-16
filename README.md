## Git Kullanım Kuralları

### Branch Oluşturma Kuralları

Yeni bir branch oluştururken aşağıdaki formata dikkat edilmelidir:

- [adın ve soyadın baş harfleri]/[yenilik veya düzeltme]/[yapılan işlemin açıklaması]-#[iş numarası veya ilgili görev]
Bu yapıda:

- **[adın ve soyadın baş harfleri]**: Branch'ı oluşturan kişinin adı ve soyadının baş harfleri. Örnek: "mustafa alp yanıkoğlu" için **may**.
- **[yenilik/düzeltme]**: Yapılan işlem bir yenilik (feature) ya da bir düzeltme (refactor) mi? Bu alan, işlemin türünü belirtir. 
- **[yapılan işlemin açıklaması]**: Kısa ve açıklayıcı bir işlem detayı. İşin ne olduğunu basitçe ifade eder.
- **#[iş numarası]**: Yapılanları daha kolay takip edebilmemiz için her işlemi numaralandırmalıyız.

#### Örnekler:
- **may/feature/login-function-added-#23**: Mustafa Alp Yanıkoğlu tarafından oluşturulan ve "giriş fonksiyonunun eklendiği" yeni bir özellik geliştirmesi.
- **bc/refactor/error-message-fixed-#56**: Başak Cengiz tarafından oluşturulan ve hata mesajlarının düzeltilmesi üzerine yapılan bir düzeltme.

### Commit Mesajı Oluşturma Kuralları

Commit mesajları oluşturulurken şu formata dikkat edilmelidir:
- [feature/refactor]: mesaj detayı
Bu yapıda:

- **feature**: Yeni bir özellik eklenmişse kullanılır.
- **refactor**: Kod düzenlemesi veya iyileştirme yapılmışsa kullanılır.
- **mesaj detayı**: Commit'in özeti, kısa ve açıklayıcı bir şekilde yazılmalıdır.

#### Örnekler:
- **feature: login işlemi için JWT entegrasyonu yapıldı**
- **refactor: kullanıcı doğrulama işlemi optimize edildi**

### Ek Notlar:
- Branch isimleri ve commit mesajları açıklayıcı, net ve kısa olmalıdır.
- Her branch, belirli bir amaca yönelik olmalıdır. Birden fazla görevi tek bir branch üzerinde yapmaktan kaçınılmalıdır.
- Commit mesajları, yapılan işin ne olduğunu net bir şekilde ifade etmelidir. Yalnızca "düzeltildi", "güncellendi" gibi genel terimler kullanmaktan kaçınılmalıdır.

Bu kurallara uymak, takım içi iş birliğini güçlendirecek ve proje yönetimini kolaylaştıracaktır.
