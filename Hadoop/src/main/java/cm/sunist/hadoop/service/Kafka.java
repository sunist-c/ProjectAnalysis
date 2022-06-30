package cm.sunist.hadoop.service;

import cm.sunist.hadoop.bean.Data;
import cm.sunist.hadoop.bean.Request;
import cm.sunist.hadoop.bean.Response;
import cm.sunist.hadoop.util.DateFormat;
import cm.sunist.hadoop.util.HiveJDBC;
import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import groovy.util.logging.Slf4j;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

import java.util.List;
import java.util.Optional;


@Component
@Slf4j
public class Kafka {
    @Autowired
    private KafkaTemplate<String,String> kafkaTemplate;

    @KafkaListener(topics = {"hive_process_request"}, groupId = "test")
    public void listen(ConsumerRecord<?, ?> record) throws Exception {
        Optional<?> kafkaMessage = Optional.ofNullable(record.value());
        if (kafkaMessage.isPresent()) {
            String message = (String) kafkaMessage.get();
            JSONObject jsonObject = JSON.parseObject(message);
            if(jsonObject.getInteger("type") == 1){
                List<Data> data = JSONArray.parseArray(jsonObject.getString("data"), Data.class);
                HiveJDBC.addData(data);
            } else if (jsonObject.getInteger("type") == 2) {
                List<Object> list = HiveJDBC.calculateCountry(DateFormat.transform(jsonObject.getString("date")), jsonObject.getInteger("data"));
                Response response = new Response(jsonObject.getInteger("data"), jsonObject.getString("uuid"), list);
                send(response);
            }else if(jsonObject.getInteger("type") == 3){
                List<Object> list = HiveJDBC.calculateProvince(DateFormat.transform(jsonObject.getString("date")), jsonObject.getInteger("data"));
                Response response = new Response(jsonObject.getInteger("data"), jsonObject.getString("uuid"), list);
                send(response);
            }else{
                System.out.println(message);
            }
        }
    }

    private void send(Response response) {
        String jsonString = JSON.toJSONString(response);
        kafkaTemplate.send("hive_process_response",jsonString);
    }
}
