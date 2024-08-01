package org.example.maven;

import net.sourceforge.pinyin4j.PinyinHelper;

public class pinyin {
    public static void main(String[] args) {
        // char c = '国';
        char c = '重';
        String[] arr = PinyinHelper.toHanyuPinyinStringArray(c);
        for (String py:arr){
            System.out.println(py);
        }
    }
}
