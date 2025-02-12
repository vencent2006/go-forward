package com.vs.myworld.business.domain;

import java.util.Date;

public class WorldEvent {
    private Long id;

    private Long worldId;

    private Integer typ;

    private String name;

    private String desc;

    private String image;

    private Date worldOccurTime;

    private Date realOccurTime;

    private Byte status;

    private Long creatorUid;

    private Date createdAt;

    private Date updatedAt;

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Long getWorldId() {
        return worldId;
    }

    public void setWorldId(Long worldId) {
        this.worldId = worldId;
    }

    public Integer getTyp() {
        return typ;
    }

    public void setTyp(Integer typ) {
        this.typ = typ;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDesc() {
        return desc;
    }

    public void setDesc(String desc) {
        this.desc = desc;
    }

    public String getImage() {
        return image;
    }

    public void setImage(String image) {
        this.image = image;
    }

    public Date getWorldOccurTime() {
        return worldOccurTime;
    }

    public void setWorldOccurTime(Date worldOccurTime) {
        this.worldOccurTime = worldOccurTime;
    }

    public Date getRealOccurTime() {
        return realOccurTime;
    }

    public void setRealOccurTime(Date realOccurTime) {
        this.realOccurTime = realOccurTime;
    }

    public Byte getStatus() {
        return status;
    }

    public void setStatus(Byte status) {
        this.status = status;
    }

    public Long getCreatorUid() {
        return creatorUid;
    }

    public void setCreatorUid(Long creatorUid) {
        this.creatorUid = creatorUid;
    }

    public Date getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(Date createdAt) {
        this.createdAt = createdAt;
    }

    public Date getUpdatedAt() {
        return updatedAt;
    }

    public void setUpdatedAt(Date updatedAt) {
        this.updatedAt = updatedAt;
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append(getClass().getSimpleName());
        sb.append(" [");
        sb.append("Hash = ").append(hashCode());
        sb.append(", id=").append(id);
        sb.append(", worldId=").append(worldId);
        sb.append(", typ=").append(typ);
        sb.append(", name=").append(name);
        sb.append(", desc=").append(desc);
        sb.append(", image=").append(image);
        sb.append(", worldOccurTime=").append(worldOccurTime);
        sb.append(", realOccurTime=").append(realOccurTime);
        sb.append(", status=").append(status);
        sb.append(", creatorUid=").append(creatorUid);
        sb.append(", createdAt=").append(createdAt);
        sb.append(", updatedAt=").append(updatedAt);
        sb.append("]");
        return sb.toString();
    }
}