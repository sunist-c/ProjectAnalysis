package cm.sunist.hadoop;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.kafka.annotation.EnableKafka;

@SpringBootApplication
@EnableKafka
public class HadoopApplication {

    public static void main(String[] args) {
        SpringApplication.run(HadoopApplication.class, args);
    }

}
