# pixPayment

(PT-BR) Microsserviço que irá intermediar a transação entre dois bancos. Ele recebe a transação e encaminha para o banco destino, status: pending, recebe a confirmação e muda o status para confirmed e envia a confirmação para o banco de origem, informando que o banco de destino processou o pagamento. O banco de origem nos envia uma confirmação de status completed.

(EN) Microservice that will intermediate the transaction between two banks. It receives the transaction and forwards it to the destination bank, status: pending, receives the confirmation and changes the status to confirmed and sends the confirmation to the source bank, informing that the destination bank processed the payment. The source bank sends us a completed status confirmation.

![image](https://user-images.githubusercontent.com/76974801/161656206-649ab6b5-92be-4328-9d8d-716281d630c9.png)

## Structures and layers:
- application:
   * factory: instantiate objects with many dependencies;
   * grpc: server and services available via gRPC;
   * kafka: consuming and processing transactions with Apache Kafka;
   * model: structure of objects that will receive external requests (via kafka or gRPC);
   * usecase: executes the flow of operations according to the business rules;
- cmd: registered commands to start the application and its services (CLI);
- domain: business rules;
- infrastructure:
  * db: ORM configuration and database interface;
  * repository: data persistence. Usually called by "usecases";

## Keywords:
- Hexagonal architecture / Ports and Adapters;
- Microsservice;
- React.js;
- Next.js (frontend);
- Nest.js (backend);
- Go;
- Kubernets;
- Docker;
- Kafka;
- Elastic Stack;
- gRPC;
- Fullcycle;
- Synchronous (Key consult);
- Asynchronous (Transations);
- DDD;
