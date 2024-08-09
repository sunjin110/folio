import Foundation

struct AppConfiguration {
    static let shared = AppConfiguration()
    let golioApiUrl: String
    
    private init() {
        let appConfMap = Bundle.main.infoDictionary?["AppConfig"] as! [AnyHashable: Any]
        guard let golioApiUrl = appConfMap["GolioApiUrl"] as? String else {
            self.golioApiUrl = ""
            return
        }
        self.golioApiUrl = golioApiUrl
    }
}
