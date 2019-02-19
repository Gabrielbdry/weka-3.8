package weka.helloworld;

import junit.framework.TestCase;
import junit.framework.TestSuite;

public class HelloWorldTest extends TestCase {
    /**
     * Test Hello world
     */
    public void testHelloWorld() throws Exception {
        HelloWorld hw = new HelloWorld();
        assertEquals(hw.toString(), "Hello World!");
    }

    public static TestSuite suite() {
        return new TestSuite(HelloWorldTest.class);
    }

    public static void main(String[] args) {
        junit.textui.TestRunner.run(suite());
    }
}
