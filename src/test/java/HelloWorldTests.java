import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class HelloWorldTests {
  @Test
  void sayHello() {
    assertEquals("Hello world", new HelloWorld().say());
  }
}
