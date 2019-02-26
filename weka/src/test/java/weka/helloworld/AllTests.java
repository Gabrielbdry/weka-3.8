package weka.helloworld;

import junit.framework.Test;
import junit.framework.TestSuite;
import weka.test.WekaTestSuite;

public class AllTests extends WekaTestSuite {
    /**
     * generates all the tests
     *
     * @return all the tests
     */
    public static Test suite() {
        TestSuite suite = new TestSuite();
        suite.addTest(HelloWorldTest.suite());
        return suite;
    }

    /**
     * for running the tests from commandline
     *
     * @param args the commandline arguments - ignored
     */
    public static void main(String[] args) {
        junit.textui.TestRunner.run(suite());
    }
}
