package cm.sunist.hadoop.bean;

import java.util.Objects;

public class Response {
    private Integer time;
    private String uuid;
    private Object data;

    public Response(Integer time, String uuid, Object data) {
        this.time = time;
        this.uuid = uuid;
        this.data = data;
    }

    public Response() {
    }

    @Override
    public String toString() {
        return "Response{" +
                "time=" + time +
                ", uuid='" + uuid + '\'' +
                ", data=" + data +
                '}';
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Response response = (Response) o;
        return Objects.equals(time, response.time) && Objects.equals(uuid, response.uuid) && Objects.equals(data, response.data);
    }

    @Override
    public int hashCode() {
        return Objects.hash(time, uuid, data);
    }

    public Integer getTime() {
        return time;
    }

    public void setTime(Integer time) {
        this.time = time;
    }

    public String getUuid() {
        return uuid;
    }

    public void setUuid(String uuid) {
        this.uuid = uuid;
    }

    public Object getData() {
        return data;
    }

    public void setData(Object data) {
        this.data = data;
    }
}
