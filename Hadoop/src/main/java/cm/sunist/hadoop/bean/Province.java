package cm.sunist.hadoop.bean;

import java.util.Date;
import java.util.Objects;

public class Province {
    private String countryName;
    private String provinceName;
    private Date refreshTime;
    private Integer confirm;
    private Integer death;
    private Integer recovered;

    public Province() {
    }

    public Province(String countryName, String provinceName, Date refreshTime, Integer confirm, Integer death, Integer recovered) {
        this.countryName = countryName;
        this.provinceName = provinceName;
        this.refreshTime = refreshTime;
        this.confirm = confirm;
        this.death = death;
        this.recovered = recovered;
    }

    public String getCountryName() {
        return countryName;
    }

    public void setCountryName(String countryName) {
        this.countryName = countryName;
    }

    public String getProvinceName() {
        return provinceName;
    }

    public void setProvinceName(String provinceName) {
        this.provinceName = provinceName;
    }

    public Date getRefreshTime() {
        return refreshTime;
    }

    public void setRefreshTime(Date refreshTime) {
        this.refreshTime = refreshTime;
    }

    public Integer getConfirm() {
        return confirm;
    }

    public void setConfirm(Integer confirm) {
        this.confirm = confirm;
    }

    public Integer getDeath() {
        return death;
    }

    public void setDeath(Integer death) {
        this.death = death;
    }

    public Integer getRecovered() {
        return recovered;
    }

    public void setRecovered(Integer recovered) {
        this.recovered = recovered;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Province province = (Province) o;
        return Objects.equals(countryName, province.countryName) && Objects.equals(provinceName, province.provinceName) && Objects.equals(refreshTime, province.refreshTime) && Objects.equals(confirm, province.confirm) && Objects.equals(death, province.death) && Objects.equals(recovered, province.recovered);
    }

    @Override
    public int hashCode() {
        return Objects.hash(countryName, provinceName, refreshTime, confirm, death, recovered);
    }

    @Override
    public String toString() {
        return "Province{" +
                "countryName='" + countryName + '\'' +
                ", provinceName='" + provinceName + '\'' +
                ", refreshTime=" + refreshTime +
                ", confirm=" + confirm +
                ", death=" + death +
                ", recovered=" + recovered +
                '}';
    }
}
