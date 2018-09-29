# Mail-server

Prerequisites:   
 1.Knowledge on SMTP (Kurose and Ross, RFC for SMTP)  

Basis : C-S architecture. 

Objectives:
  1. DB under the application to maintain the e-mails.
  2. HTTP to view emails.
  3. Server: DNS resolution, DB maintaining, mail sending recieving.
  4. DB : Capability to maintain all sent and inbox. 
  5. Service to view my account. 

Language : C,C++


Resources used:  
  1. RFC 821, SMTP protocol, https://tools.ietf.org/html/rfc821  
  2. The latest RFC for SMTP protocol, RFC 5321, https://tools.ietf.org/html/rfc5321  (This will be used as it is the latest) 
  3. Kurose and Ross for SMTP basics https://bit.ly/2QeCuhl
  4. Companion to RFC 5321 is RFC 5322 https://tools.ietf.org/html/rfc5322
  

About the protocol:
  1. RFC 5321 discusses transfer of mail over TCP.
  2. A network consists of mutually-TCP-accessible hosts on the public internet. Using SMTP one process can transfer mail to       another process via a gateway or relay process which is accessible to both networks.
  3. The mail may pass through many intermediate hosts and reach the ultimate recipient.
  4. The basic design is as follows - 
  
  <pre>
                                  +----------+                +----------+  
                      +------+    |          |                |          |  
                      | User |<-->|          |      SMTP      |          |  
                      +------+    |  Client- |Commands/Replies| Server-  |  
                      +------+    |   SMTP   |<-------------->|    SMTP  |    +------+  
                      | File |<-->|          |    and Mail    |          |<-->| File |  
                      |System|    |          |                |          |    |System|  
                      +------+    +----------+                +----------+    +------+  
                                  SMTP client                SMTP server  
                                
      
  </pre>
                     
   5. The client and server have a two way transmission channel established between them.
   6. The responsibility of an SMTP client is to transfer mail messages to one or more SMTP servers,or report its failure to do so.
   7. An SMTP client determines the address of an appropriate host running an SMTP server by resolving a destination domain name to either an intermediate Mail eXchanger host or a final target host.
   8. SMTP *commands* are generated by the SMTP client and sent to the SMTP server. SMTP *replies* are sent from the server to the client in response to these commands.The replies may indicate command was accepted, more commands are expected or temporary or permanent error exists.
   9. Once the server has issued a success response at the end of the mail data, a formal handoff of responsibility for the message occurs: the protocol requires that a server MUST accept responsibility for either delivering the message or properly reporting the failure to do so.
   10. After the transmission channel is established, there will be a series of commands to specify the origin or the mail, the recipient and about the transmission. 
   11. After the transmission is complete, the client may request to shut down the connection or may initiate other mail transactions.
   12. 
