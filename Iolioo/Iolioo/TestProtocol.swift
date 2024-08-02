//
//  TestProtocol.swift
//  Iolioo
//
//  Created by 尹舜真 on 2024/08/01.
//

import Foundation
import Spyable

@Spyable
public protocol SampleRepository {
    func get(id: String) -> Int
}
