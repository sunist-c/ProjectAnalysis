package cm.sunist.hadoop.bean;

import java.util.Objects;

public class Request {
    private String uuid;
    private String date;    //指数据的日期，即爬虫中的refresh_time部分
    private Integer type;   //请求类型，1表示导入爬虫的数据，2表示计算国家数据，3表示计算省级数据
    private Object data;    //如果type = 1，data为爬虫的数据
                            //如果type = 2 或者 3，data是要计算的时间段长度

    public Request() {
    }

    public Request(String uuid, String date, Integer type, Object data) {
        this.uuid = uuid;
        this.date = date;
        this.type = type;
        this.data = data;
    }

    @Override
    public String toString() {
        return "Request{" +
                "uuid='" + uuid + '\'' +
                ", date='" + date + '\'' +
                ", type=" + type +
                ", data=" + data +
                '}';
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Request request = (Request) o;
        return Objects.equals(uuid, request.uuid) && Objects.equals(date, request.date) && Objects.equals(type, request.type) && Objects.equals(data, request.data);
    }

    @Override
    public int hashCode() {
        return Objects.hash(uuid, date, type, data);
    }

    public String getUuid() {
        return uuid;
    }

    public void setUuid(String uuid) {
        this.uuid = uuid;
    }

    public String getDate() {
        return date;
    }

    public void setDate(String date) {
        this.date = date;
    }

    public Integer getType() {
        return type;
    }

    public void setType(Integer type) {
        this.type = type;
    }

    public Object getData() {
        return data;
    }

    public void setData(Object data) {
        this.data = data;
    }
}
