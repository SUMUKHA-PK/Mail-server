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
  

Summary of the protocol:
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
   8. SMTP *commands* are generated by the SMTP client and sent to the SMTP server.All commands begin with a *command verb* (this is not case senstitive). SMTP *replies* are sent from the server to the client in response to these commands.The replies may indicate command was accepted, more commands are expected or temporary or permanent error exists.(Permanent = same message again will result in same error, Temporary = retransmission after a while will solve problem) Replies are acknowledgements(Positive or negative) generally numeric completion codes, followed by a text string.  
   9. Once the server has issued a success response at the end of the mail data, a formal handoff of responsibility for the message occurs: the protocol requires that a server MUST accept responsibility for either delivering the message or properly reporting the failure to do so.
   10. After the transmission channel is established, there will be a series of commands to specify the origin or the mail, the recipient and about the transmission. 
   11. After the transmission is complete, the client may request to shut down the connection or may initiate other mail transactions.
   12. SMTP transfers a *mail object*. A mail object contains an envelope and content. The envelope (SMTP envelope) is sent as a series of SMTP protocol units , including originator address, one or more recipient address and optional protocol extension material.
   13. The local part of the mail is case sensitive (the sender and reciever addresses).
   14. In general, a relay SMTP SHOULD assume that the message content it has received is valid and, assuming that the envelope permits doing so, relay it without inspecting that content.  Of course, if the content is mislabeled and the data path cannot accept the actual content, this may result in the ultimate delivery of a severely garbled message to the recipient. Delivery SMTP systems MAY reject such messages, or return them as undeliverable, rather than deliver them.  In the absence of a server- offered extension explicitly permitting it, a sending SMTP system is not permitted to send envelope commands in any character set other than US-ASCII.  Receiving systems SHOULD reject such commands,normally using "500 syntax error - invalid character" replies.
   15. _Initiation of a session_ is said to occur when a client opens a connection with the  SMTP server and server replies with an _opening message_. This message may contain a version and software specification of the server.
   16. SMTP protocol allows the server to formally reject the connection request by sending a 554 code instead of a 220. But, the server must still wait for the client to send a QUIT command. All sequences of commands in between must be responded with a _bad sequence of commands_ since connection is not established. The server returing this reply must respond with enough information in the replies enough to debug the situation that caused the error in the client.
   17. Once the server sends the greeting,the client responds with a EHLO command to the server, which indicates the clients identity. EHLO indicates the client is able to process and requests the server to do the same. 
   18. The MAIL command : MAIL FROM:<reverse-path> [SP <mail-parameters> ] <CRLF>   --> 250 OK when accepted
This command tells the SMTP-receiver that a new mail transaction is starting and to reset all its buffers and state tables, recipientsand mail data. 
   19. The RCPT command : RCPT TO:<forward-path> [ SP <rcpt-parameters> ] <CRLF>    --> 250 OK when accepted
This step of the procedure can be repeated any number of times. Server stores the forward path once the command is accepted. If the forward-path is not a deliverable address, a 550 reply is generated,saying "no such user exists". The forward-path can contain more than one mailbox.
   20. The DATA command : DATA <CRLF>   -->354 intermediate reply and a 250 OK reply after reading completely.
Since the mail data is sent on the transmission channel, the end of mail data must be indicated so that the command and reply dialog can be resumed.  SMTP indicates the end of the mail data by sending a line containing only a "." (period or full stop).
   If there was no MAIL, or no RCPT, command, or all such commands were rejected, the server MAY return a "command out of sequence" (503) or    "no valid recipients" (554) reply in response to the DATA command. 
   21. While forwarding, one of the two following things may happen. One, the server may silently forward and return a 250 OK or 251 when there is an address change.Two, reject with 551 saying address updating is necessary or 550 with no address-specific information.
   22. The _verify command_ : (VRFY) Verifies the user name(obtained as string input), if normal, a 250 response is given, 553 for ambiguos(user ambigous message).
      <br>  Eg: 
      <br>  553- Ambiguous; Possibilities are
      <br>  553-Joe Smith <jsmith@foo.com>
      <br>  553-Harry Smith <hsmith@foo.com>
      <br>  553 Melvin Smith <dweep@foo.com>
   23. The EXPN command : Used to obtain the content of the mailing list. Here the input string identifies a mailing list and successful replies (250) may include full name of the users and must include mailboxes on the mailing list.  
       <br> Eg: The case of expanding a mailbox list requires a multiline reply, such as:
       <br> C: EXPN Example-People
       <br> S: 250-Jon Postel <Postel@isi.edu>
       <br> S: 250-Fred Fonebone <Fonebone@physics.foo-u.edu>
       <br> S: 250 Sam Q. Smith <SQSmith@specific.generic.com>
                       or
       <br> C: EXPN Executive-Washroom-List
       <br> S: 550 Access Denied to You.
   24. A server MUST NOT return 250 if all it has done is to verify that the syntax given is valid.  In that case, 502 (Command not implemented) or 500 (Syntax error, command unrecognized) SHOULD be returned. (For VRFY and EXPN)
   25.  
