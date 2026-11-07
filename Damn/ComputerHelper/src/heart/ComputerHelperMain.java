package heart;
import java.io.PrintStream;
import java.nio.charset.StandardCharsets;
import java.util.Scanner;

import heart.libs.TestingClass;

public class ComputerHelperMain {
    public static void main(String args[]) {
        System.setOut(new PrintStream(System.out, true, StandardCharsets.UTF_8));
        System.setErr(new PrintStream(System.err, true, StandardCharsets.UTF_8));
        TestingClass testing = new TestingClass();
        Scanner sc = new Scanner(System.in);

        testing.Test();
        sc.nextLine();
    }
}