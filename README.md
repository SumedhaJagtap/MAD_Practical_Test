# MAD_Practical_Test
Name : Sumedha Jagtap
Roll No : 16218
Date : 7 April 2019


1. There will be n Channels for each cashier.

2. Number of Cashier,Customer and TimePerCustomer will be taken from command-line.

3. var mux = sync.Mutex{} 
   Mutex lock will be declared globally

4. 
    4.1
        func allocateCust(custNo int, cashierNo int, timePerCustomer int, cashier chan<- int){
        /*
        This function will take the customer's number in a queue
        Put that number into a sender channer.
        Then will print a msg Cashier 'n': Customer custNo  Started
        Later receiving channel will consume it.
        Before sending the Mutex will acquire a lock on this.
        Later Lock will be Unlocked.
        */

        mux.Lock()
        cashier <- custNo
        fmt.Printf("Cashier %d: Customer %d Started", cashierNo, custNo)
        time.Sleep(timePerCustomer * time.Second)
        mux.Unlock()
        
        
        }
    4.2



5. In main :

    5.1  n number of channels will be declared
   
    5.2  waitGroup will be created
    
    5.3  wg.Add(c)
            - 'c' a number of cashiers will be added into a waitgroup.
    
    5.4 Goroutine will be used to call to allocateCust()
       5.4.1  For loop of n number of customers will be iterated with a call to a function
              allocateCust()
       5.4.2  A goroutine will be evaluted with ith customer to call a allocateCust(). 

        /* For now I'm confused in how to assign a function call for each cashier 
        Please ignore the bank.go code*/
        
         for i:=1;i<=CustNo;i++{
             go func(i int){
                defer wg.Done()
                if _,ok:=<-c1;!ok{
                
                    allocateCust(i,1,timePerSecond,c1)
                
                } else if _,ok:=<-c2;!ok{
                
                    allocateCust(i,2,timePerSecond,c1)
                
                } else if _,ok:=<-c3;!ok{
                
                    allocateCust(i,3,timePerSecond,c1)
                
                } 
             }(i)
         }
         wg.Wait()

        5.4.2 All the cashier channels with be emptyed again by looping over all channels
                for i:=1;i<=100;i++{
                    select{
                        case msg1:=<-c1:
                            fmt.Printf("Cashier 1: Customer %d Completed", i)
                         case msg2:=<-c2:
                            fmt.Printf("Cashier 2: Customer %d Completed", i)
                         case msg3:=<-c3:
                            fmt.Printf("Cashier 3: Customer %d Completed", i)
                    }
                }
                c1.Close()
                c2.Close()
                c3.Close()