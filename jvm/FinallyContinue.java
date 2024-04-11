/**
 * Classfile jvm/FinallyContinue.class
 * Last modified 2024-4-11; size 767 bytes
 * MD5 checksum b30800f8d830bcc7b34463565728cf6a
 * Compiled from "FinallyContinue.java"
 * public class FinallyContinue
 * minor version: 0
 * major version: 52
 * flags: ACC_PUBLIC, ACC_SUPER
 * Constant pool:
 * #1 = Methodref #12.#24 // java/lang/Object."<init>":()V
 * #2 = Long 100l
 * #4 = Methodref #25.#26 // java/lang/Thread.sleep:(J)V
 * #5 = Fieldref #27.#28 // java/lang/System.out:Ljava/io/PrintStream;
 * #6 = Methodref #29.#30 // java/io/PrintStream.println:(I)V
 * #7 = String #31 // Finally
 * #8 = Methodref #29.#32 // java/io/PrintStream.println:(Ljava/lang/String;)V
 * #9 = Class #33 // java/lang/Exception
 * #10 = String #34 // Exception
 * #11 = Class #35 // FinallyContinue
 * #12 = Class #36 // java/lang/Object
 * #13 = Utf8 <init>
 * #14 = Utf8 ()V
 * #15 = Utf8 Code
 * #16 = Utf8 LineNumberTable
 * #17 = Utf8 main
 * #18 = Utf8 ([Ljava/lang/String;)V
 * #19 = Utf8 StackMapTable
 * #20 = Class #33 // java/lang/Exception
 * #21 = Class #37 // java/lang/Throwable
 * #22 = Utf8 SourceFile
 * #23 = Utf8 FinallyContinue.java
 * #24 = NameAndType #13:#14 // "<init>":()V
 * #25 = Class #38 // java/lang/Thread
 * #26 = NameAndType #39:#40 // sleep:(J)V
 * #27 = Class #41 // java/lang/System
 * #28 = NameAndType #42:#43 // out:Ljava/io/PrintStream;
 * #29 = Class #44 // java/io/PrintStream
 * #30 = NameAndType #45:#46 // println:(I)V
 * #31 = Utf8 Finally
 * #32 = NameAndType #45:#47 // println:(Ljava/lang/String;)V
 * #33 = Utf8 java/lang/Exception
 * #34 = Utf8 Exception
 * #35 = Utf8 FinallyContinue
 * #36 = Utf8 java/lang/Object
 * #37 = Utf8 java/lang/Throwable
 * #38 = Utf8 java/lang/Thread
 * #39 = Utf8 sleep
 * #40 = Utf8 (J)V
 * #41 = Utf8 java/lang/System
 * #42 = Utf8 out
 * #43 = Utf8 Ljava/io/PrintStream;
 * #44 = Utf8 java/io/PrintStream
 * #45 = Utf8 println
 * #46 = Utf8 (I)V
 * #47 = Utf8 (Ljava/lang/String;)V
 * {
 * public FinallyContinue();
 * descriptor: ()V
 * flags: ACC_PUBLIC
 * Code:
 * stack=1, locals=1, args_size=1
 * 0: aload_0
 * 1: invokespecial #1 // Method java/lang/Object."<init>":()V
 * 4: return
 * LineNumberTable:
 * line 1: 0
 * 
 * public static void main(java.lang.String[]);
 * descriptor: ([Ljava/lang/String;)V
 * flags: ACC_PUBLIC, ACC_STATIC
 * Code:
 * stack=3, locals=5, args_size=1
 * 0: iconst_0
 * 1: istore_1
 * 2: iconst_5
 * 3: istore_2
 * 4: iload_1
 * 5: iinc 1, 1
 * 8: iload_2
 * 9: if_icmpge 79
 * 12: iload_1
 * 13: iload_2
 * 14: iconst_2
 * 15: idiv
 * 16: if_icmpne 25
 * 19: ldc2_w #2 // long 100l
 * 22: invokestatic #4 // Method java/lang/Thread.sleep:(J)V
 * 25: getstatic #5 // Field java/lang/System.out:Ljava/io/PrintStream;
 * 28: iload_1
 * 29: invokevirtual #6 // Method java/io/PrintStream.println:(I)V
 * 32: getstatic #5 // Field java/lang/System.out:Ljava/io/PrintStream;
 * 35: ldc #7 // String Finally
 * 37: invokevirtual #8 // Method
 * java/io/PrintStream.println:(Ljava/lang/String;)V
 * 40: goto 76
 * 43: astore_3
 * 44: getstatic #5 // Field java/lang/System.out:Ljava/io/PrintStream;
 * 47: ldc #10 // String Exception
 * 49: invokevirtual #8 // Method
 * java/io/PrintStream.println:(Ljava/lang/String;)V
 * 52: getstatic #5 // Field java/lang/System.out:Ljava/io/PrintStream;
 * 55: ldc #7 // String Finally
 * 57: invokevirtual #8 // Method
 * java/io/PrintStream.println:(Ljava/lang/String;)V
 * 60: goto 76
 * 63: astore 4
 * 65: getstatic #5 // Field java/lang/System.out:Ljava/io/PrintStream;
 * 68: ldc #7 // String Finally
 * 70: invokevirtual #8 // Method
 * java/io/PrintStream.println:(Ljava/lang/String;)V
 * 73: aload 4
 * 75: athrow
 * 76: goto 4
 * 79: return
 * Exception table:
 * from to target type
 * 12 32 43 Class java/lang/Exception
 * 12 32 63 any
 * 43 52 63 any
 * 63 65 63 any
 * LineNumberTable:
 * line 3: 0
 * line 4: 2
 * line 5: 4
 * line 7: 12
 * line 8: 19
 * line 10: 25
 * line 16: 32
 * line 17: 40
 * line 12: 43
 * line 13: 44
 * line 16: 52
 * line 17: 60
 * line 16: 63
 * line 17: 73
 * line 19: 79
 * StackMapTable: number_of_entries = 6
 * frame_type = 253
 * offset_delta = 4
 * locals = [ int, int ]
 * frame_type = 20
 * frame_type = 81
 * stack = [ class java/lang/Exception ]
 * frame_type = 83
 * stack = [ class java/lang/Throwable ]
 * frame_type = 12
 * frame_type = 2
 * }
 * SourceFile: "FinallyContinue.java"
 */
public class FinallyContinue {
    public static void main(String[] args) {
        int i = 0;
        int n = 5;
        while (i++ < n) {
            try {
                if (i == n / 2) {
                    Thread.sleep(100);
                    ;
                }
                System.out.println(i);
            } catch (Exception e) {
                System.out.println("Exception");
            } finally {
                System.out.println("Finally");
            }
        }
    }
}
