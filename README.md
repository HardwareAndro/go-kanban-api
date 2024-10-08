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
- **feature: JWT integration was made for the login process**
- **refactor: user authentication process optimized**

## Pull Request (PR) Kuralları

### PR Süreci ve Branch Yönetimi

Doğrudan **master** veya **development** branch'lerine kod göndermek (push) kesinlikle **yasaktır**. Kodların ana branch'lara taşınması süreci belirli bir PR (Pull Request) akışı çerçevesinde gerçekleştirilmelidir. Bu süreç şu şekilde işlemelidir:

#### 1. Kendi Branch'ını Aç
Her geliştirici, yapacağı işle ilgili olarak kendi adına yeni bir branch açmalıdır. Branch isimlendirme kurallarına dikkat edilmelidir (örneğin: **may/feature/login-function-added-#23**).

#### 2. Development Branch'ine PR Aç
Branch'te yapılan geliştirmeler tamamlandığında ve testler başarıyla geçtiğinde, kodların doğrudan **development** branch'ine merge edilmesi için bir **Pull Request (PR)** açılmalıdır. Bu PR'de şu süreçlere dikkat edilmelidir:

- Yapılan işin açıklayıcı ve net bir başlığı olmalıdır.
- Hangi özelliklerin eklendiği ya da hangi düzeltmelerin yapıldığı detaylı bir şekilde açıklanmalıdır.
- Yapılan geliştirmelerle ilgili testler çalıştırılmalı ve test sonuçları paylaşılmalıdır (varsa otomatik test entegrasyonları).
- Kod incelemesi (**code review**) sürecini başlatmak için ilgili takım üyeleri PR'ye atanmalıdır.

#### 3. Development Branch'inde Sorunsuz Merge
Kodlar inceleme sonrası onaylandığında ve gerekli düzenlemeler yapıldıktan sonra PR, **development** branch'ine merge edilir.

#### 4. Master Branch'ine PR Aç
Geliştirmelerin **development** branch'inde stabil olduğundan emin olduktan sonra, **master** branch'ine merge edilmek üzere bir PR açılır. Bu PR'nin süreci de yukarıda belirtilen adımlarla aynı olmalıdır.

- **Master** branch'ine yalnızca **development** branch'i üzerinden PR açılabilir.
- **Master** branch'ine yapılacak her PR, proje için önemli bir adım olduğundan daha dikkatli incelenmelidir.

---

### Neden Doğrudan Master veya Development Branch'lerine Push Yapılmamalı?

- **Kod Kalitesini Arttırır:** PR süreci, kodun gözden geçirilmesini sağlar ve böylece hataların veya iyileştirilebilir noktaların tespit edilmesine yardımcı olur.
- **Takım İçinde Şeffaflık Sağlar:** PR süreci ile hangi işin ne zaman, kim tarafından yapıldığı daha net bir şekilde takip edilebilir.
- **Hataları Erken Yakalar:** Development branch'inde merge edilen kodlar, topluca **master** branch'ine taşınmadan önce test edilip stabilize edilebilir. Bu sayede **master** branch'ine daha güvenilir ve sağlam kodlar gönderilir.


### Ek Notlar:
- Branch isimleri ve commit mesajları açıklayıcı, net ve kısa olmalıdır.
- Her branch, belirli bir amaca yönelik olmalıdır. Birden fazla görevi tek bir branch üzerinde yapmaktan kaçınılmalıdır.
- Commit mesajları, yapılan işin ne olduğunu net bir şekilde ifade etmelidir. Yalnızca "düzeltildi", "güncellendi" gibi genel terimler kullanmaktan kaçınılmalıdır.

Bu kurallara uymak, takım içi iş birliğini güçlendirecek ve proje yönetimini kolaylaştıracaktır.
