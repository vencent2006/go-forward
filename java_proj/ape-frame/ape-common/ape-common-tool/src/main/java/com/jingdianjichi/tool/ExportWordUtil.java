package com.jingdianjichi.tool;

import freemarker.template.Configuration;
import freemarker.template.Template;

import javax.servlet.ServletOutputStream;
import javax.servlet.http.HttpServletResponse;
import java.io.*;
import java.net.URLEncoder;
import java.util.Map;

/**
 * @Author: ChickenWing
 * @Description: 导出word的util
 * @DateTime: 2022/10/23 22:01
 */
public class ExportWordUtil {

    private static Configuration configuration = null;

    private static final String SUFFIX = ".doc";

    static {
        configuration = new Configuration();
        configuration.setDefaultEncoding("UTF-8");
        configuration.setClassForTemplateLoading(ExportWordUtil.class, "/template/word");
    }

    public static void exportWord(Map map, String title, String ftlName) throws Exception {
        Template template = configuration.getTemplate(ftlName);
        File file = null;
        InputStream inputStream = null;
        ServletOutputStream out = null;
        file = createDocFile(map, template);
        inputStream = new FileInputStream(file);
        String fileName = title + SUFFIX;
        HttpServletResponse response = SpringContextUtils.getHttpServletResponse();
        response.setCharacterEncoding("utf-8");
        response.setContentType("application/msword");
        response.setHeader("Content-Disposition", "attachment;filename=" + URLEncoder.encode(fileName, "utf-8"));
        out = response.getOutputStream();
        byte[] buffer = new byte[512];
        int bytesToRead = -1;
        while ((bytesToRead = inputStream.read(buffer)) != -1) {
            out.write(buffer, 0, bytesToRead);
        }
        if (inputStream != null) {
            inputStream.close();
        }
        if (out != null) {
            out.close();
        }
        if (file != null) {
            file.delete();
        }
    }

    private static File createDocFile(Map dataMap, Template template) throws Exception {
        File file = new File("init.doc");
        Writer writer = new OutputStreamWriter(new FileOutputStream(file), "utf-8");
        template.process(dataMap, writer);
        writer.close();
        return file;
    }


}
