//
//  LoginView.swift
//  molio
//
//  Created by 尹舜真 on 2024/07/02.
//

import SwiftUI

struct LoginView: View {
    var body: some View {
        Text("Login").font(.title).padding()
        Button(action: signIn) {
            Text("Google Login")
        }.padding()
    }
    
    private func signIn() {
        print("sign in")
    }
}

#Preview {
    LoginView()
}
