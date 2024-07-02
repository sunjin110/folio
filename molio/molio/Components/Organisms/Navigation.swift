//
//  Navigation.swift
//  molio
//
//  Created by 尹舜真 on 2024/07/02.
//

import SwiftUI

struct Navigation: View {
    
    public let contents: [Content]
    
    var body: some View {
        NavigationStack {
           
        }
    }
}

class Content {
    public let title: String
    
    init(title: String) {
        self.title = title
    }
}

#Preview {
    
    let contents: [Content] = [Content(title: "test")]
    
    Navigation(contents: contents)
}
