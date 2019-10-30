# 実装の手順

* domain層->domain層には技術的関心事を持ち込まない

1. domain層ーmodel
    
    →modelでは構造体の定義

2. domain層-repository
    
    →repositoryではDBやKVSなどで行うCRUD処理を担う
    
    →interfaceを定義
