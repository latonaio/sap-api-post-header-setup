# sap-api-post-header-setup  
sap-api-post-header-setup は、主にエッジコンピューティング環境において、[sap-sandbox](https://github.com/latonaio/sap-sandbox) が対象とする SAP APIs ならびに 各 SAP API Integrations の Runtimes のうち「creates」に分類される APIs ならびに Runtimes について、当該 Runtimes において Post する Json フォーマットに対して必要なヘッダ情報を付与するマイクロサービスです。

## 動作環境  
sap-api-post-header-setup は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  

* エッジ Kubernetes （推奨）  
* AION のリソース （推奨)  
* OS: LinuxOS （必須）  
* CPU: ARM/AMD/Intel（いずれか必須）  
* Golang Runtime 

## クラウド環境での利用
sap-api-post-header-setup は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## JSON に 付与される ヘッダデータ
sap-api-post-header-setup では、次のヘッダデータが付与されます。

* Accept値（固定値 "application/json"）
* Content-Type値（固定値 "application/json"）
* csrfToken（コール先のSAPサーバーの csrfToken）
* refreshTokenURL（コール先のSAPサーバーの csrfTokenURL）
* SAPUser（SAPユーザID）
* Pass（SAPパスワード）

## go.mod / go.sum
sap-api-post-header-setup は、ライブラリであり、go.mod / go.sum に設定することで、他のレポジトリやランタイムで実行できます。  
sap-api-post-header-setup は、[sap-sandbox](https://github.com/latonaio/sap-sandbox)における SAP APIs ならびに 各 SAP API Integrations の Runtimes を対象としています。  
sap-api-post-header-setup は、マイクロサービスとして利用されることができます。  
