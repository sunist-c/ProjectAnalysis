package cn.sunist.server.service;

import cn.sunist.server.bean.Data;
import cn.sunist.server.bean.Response;
import cn.sunist.server.util.DateFormat;
import cn.sunist.server.dao.HiveJDBC;
import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import java.util.List;
import java.util.Optional;


@Component
public class Kafka {
    @Resource
    private KafkaTemplate<String,String> kafkaTemplate;

    @KafkaListener(topics = {"hive_process_request"},groupId = "test")
    public void listen(ConsumerRecord<?, ?> record){
        Optional<?> kafkaMessage = Optional.ofNullable(record.value());
        if (kafkaMessage.isPresent()) {
            try{
                String message = (String) kafkaMessage.get();
                System.out.println(message);
                JSONObject jsonObject = JSON.parseObject(message);
                if(jsonObject.getInteger("type") == 1){
                    List<Data> data = JSONArray.parseArray(jsonObject.getString("data"), Data.class);
                    HiveJDBC.addData(data);
                    Response response = new Response(jsonObject.getInteger("data"), jsonObject.getString("uuid"), "");
                    send(response);
                } else if (jsonObject.getInteger("type") == 2) {
                    List<Object> list = HiveJDBC.calculateCountry(DateFormat.transform(jsonObject.getString("date")), jsonObject.getInteger("data"));
                    Response response = new Response(jsonObject.getInteger("data"), jsonObject.getString("uuid"), list);
                    send(response);
                }else if(jsonObject.getInteger("type") == 3){
                    List<Object> list = HiveJDBC.calculateProvince(DateFormat.transform(jsonObject.getString("date")), jsonObject.getInteger("data"));
                    Response response = new Response(jsonObject.getInteger("data"), jsonObject.getString("uuid"), list);
                    send(response);
                }
            }catch (Exception e){
                e.printStackTrace();
            }
        }
    }

    private void send(Response response) {
        String jsonString = JSON.toJSONString(response);
        System.out.println(jsonString);
        kafkaTemplate.send("hive_process_response",jsonString);
    }
}
