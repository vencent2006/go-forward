package com.vs.myworld.business.util;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

public class PrintUtil {
    public static void printJsonFormat(Object src) {
        printSepLine();
        Gson gson = new GsonBuilder().setPrettyPrinting().create();
        String json = gson.toJson(src);
        System.out.println(">>>> " + Thread.currentThread().getStackTrace()[2].getMethodName());
        System.out.println(json);
    }

    static void printSepLine() {
        System.out.println("*          *   *   *         *       ");
        System.out.println(" *        *   *   * *       *        ");
        System.out.println("  *      *   *   *   *     *         ");
        System.out.println("   *    *   *   *     *   *          ");
        System.out.println("    *  *   *   *       * *           ");
        System.out.println("     *    *   *         *            ");
        System.out.println();
    }
}
